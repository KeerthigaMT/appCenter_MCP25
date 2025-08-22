package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/app-center-client/mcp-server/config"
	"github.com/app-center-client/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Orginvitations_sendnewinvitationHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		org_nameVal, ok := args["org_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: org_name"), nil
		}
		org_name, ok := org_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: org_name"), nil
		}
		emailVal, ok := args["email"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: email"), nil
		}
		email, ok := emailVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: email"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/v0.1/orgs/%s/invitations/%s/resend", cfg.BaseURL, org_name, email)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("X-API-Token", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateOrginvitations_sendnewinvitationTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v0_1_orgs_org_name_invitations_email_resend",
		mcp.WithDescription("Cancels an existing organization invitation for the user and sends a new one"),
		mcp.WithString("org_name", mcp.Required(), mcp.Description("The organization's name")),
		mcp.WithString("email", mcp.Required(), mcp.Description("The email address of the user to send the password reset mail to.")),
		mcp.WithString("role", mcp.Description("Input parameter: The role of the user to be added")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Orginvitations_sendnewinvitationHandler(cfg),
	}
}

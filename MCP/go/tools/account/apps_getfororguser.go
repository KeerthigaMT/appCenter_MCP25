package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/app-center-client/mcp-server/config"
	"github.com/app-center-client/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Apps_getfororguserHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		user_nameVal, ok := args["user_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: user_name"), nil
		}
		user_name, ok := user_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: user_name"), nil
		}
		url := fmt.Sprintf("%s/v0.1/orgs/%s/users/%s/apps", cfg.BaseURL, org_name, user_name)
		req, err := http.NewRequest("GET", url, nil)
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
		var result []interface{}
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

func CreateApps_getfororguserTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v0_1_orgs_org_name_users_user_name_apps",
		mcp.WithDescription("Get a user apps information from an organization by name"),
		mcp.WithString("org_name", mcp.Required(), mcp.Description("The organization's name")),
		mcp.WithString("user_name", mcp.Required(), mcp.Description("The slug name of the user")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Apps_getfororguserHandler(cfg),
	}
}

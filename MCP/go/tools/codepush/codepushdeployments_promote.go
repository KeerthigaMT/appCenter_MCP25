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

func Codepushdeployments_promoteHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		deployment_nameVal, ok := args["deployment_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: deployment_name"), nil
		}
		deployment_name, ok := deployment_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: deployment_name"), nil
		}
		promote_deployment_nameVal, ok := args["promote_deployment_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: promote_deployment_name"), nil
		}
		promote_deployment_name, ok := promote_deployment_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: promote_deployment_name"), nil
		}
		owner_nameVal, ok := args["owner_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: owner_name"), nil
		}
		owner_name, ok := owner_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: owner_name"), nil
		}
		app_nameVal, ok := args["app_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: app_name"), nil
		}
		app_name, ok := app_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: app_name"), nil
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
		url := fmt.Sprintf("%s/v0.1/apps/%s/%s/deployments/%s/promote_release/%s", cfg.BaseURL, deployment_name, promote_deployment_name, owner_name, app_name)
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
		var result interface{}
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

func CreateCodepushdeployments_promoteTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v0_1_apps_owner_name_app_name_deployments_deployment_name_promote_release_promote_deployment_name",
		mcp.WithDescription("Promote one release (default latest one) from one deployment to another"),
		mcp.WithString("deployment_name", mcp.Required(), mcp.Description("deployment name")),
		mcp.WithString("promote_deployment_name", mcp.Required(), mcp.Description("deployment name")),
		mcp.WithString("owner_name", mcp.Required(), mcp.Description("The name of the owner")),
		mcp.WithString("app_name", mcp.Required(), mcp.Description("The name of the application")),
		mcp.WithString("description", mcp.Description("")),
		mcp.WithBoolean("is_disabled", mcp.Description("")),
		mcp.WithBoolean("is_mandatory", mcp.Description("")),
		mcp.WithNumber("rollout", mcp.Description("")),
		mcp.WithString("target_binary_range", mcp.Description("")),
		mcp.WithString("label", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Codepushdeployments_promoteHandler(cfg),
	}
}

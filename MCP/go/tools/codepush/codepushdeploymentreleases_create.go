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

func Codepushdeploymentreleases_createHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/v0.1/apps/%s/%s/deployments/%s/releases", cfg.BaseURL, deployment_name, owner_name, app_name)
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

func CreateCodepushdeploymentreleases_createTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v0_1_apps_owner_name_app_name_deployments_deployment_name_releases",
		mcp.WithDescription("Create a new CodePush release for the specified deployment"),
		mcp.WithString("deployment_name", mcp.Required(), mcp.Description("deployment name")),
		mcp.WithString("owner_name", mcp.Required(), mcp.Description("The name of the owner")),
		mcp.WithString("app_name", mcp.Required(), mcp.Description("The name of the application")),
		mcp.WithString("release_upload", mcp.Required(), mcp.Description("Input parameter: The upload metadata from the release initialization step.")),
		mcp.WithNumber("rollout", mcp.Description("Input parameter: This specifies the percentage of users (as an integer between 1 and 100) that should be eligible to receive this update.")),
		mcp.WithString("target_binary_version", mcp.Required(), mcp.Description("Input parameter: the binary version of the application")),
		mcp.WithString("deployment_name", mcp.Description("Input parameter: This specifies which deployment you want to release the update to. Default is Staging.")),
		mcp.WithString("description", mcp.Description("Input parameter: This provides an optional \"change log\" for the deployment.")),
		mcp.WithBoolean("disabled", mcp.Description("Input parameter: This specifies whether an update should be downloadable by end users or not.")),
		mcp.WithBoolean("mandatory", mcp.Description("Input parameter: This specifies whether the update should be considered mandatory or not (e.g. it includes a critical security fix).")),
		mcp.WithBoolean("no_duplicate_release_error", mcp.Description("Input parameter: This specifies that if the update is identical to the latest release on the deployment, the CLI should generate a warning instead of an error.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Codepushdeploymentreleases_createHandler(cfg),
	}
}

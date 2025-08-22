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

func Releases_getlatestpublicreleaseHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		app_secretVal, ok := args["app_secret"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: app_secret"), nil
		}
		app_secret, ok := app_secretVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: app_secret"), nil
		}
		url := fmt.Sprintf("%s/v0.1/public/sdk/apps/%s/releases/latest", cfg.BaseURL, app_secret)
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

func CreateReleases_getlatestpublicreleaseTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v0_1_public_sdk_apps_app_secret_releases_latest",
		mcp.WithDescription("Get the latest public release for the given app."),
		mcp.WithString("app_secret", mcp.Required(), mcp.Description("The secret of the target application")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Releases_getlatestpublicreleaseHandler(cfg),
	}
}

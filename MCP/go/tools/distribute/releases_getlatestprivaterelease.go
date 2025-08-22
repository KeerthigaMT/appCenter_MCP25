package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/app-center-client/mcp-server/config"
	"github.com/app-center-client/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Releases_getlatestprivatereleaseHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["udid"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("udid=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v0.1/sdk/apps/%s/releases/private/latest%s", cfg.BaseURL, app_secret, queryString)
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

func CreateReleases_getlatestprivatereleaseTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v0_1_sdk_apps_app_secret_releases_private_latest",
		mcp.WithDescription("Get the latest release distributed to a private group the given user is a member of for the given app."),
		mcp.WithString("app_secret", mcp.Required(), mcp.Description("The secret of the target application")),
		mcp.WithString("udid", mcp.Description("When passing `udid` in the query string, a provisioning check for the given device ID will be done. Will be ignored for non-iOS platforms.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Releases_getlatestprivatereleaseHandler(cfg),
	}
}

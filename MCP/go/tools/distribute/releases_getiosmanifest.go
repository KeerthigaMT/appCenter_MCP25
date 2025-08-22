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

func Releases_getiosmanifestHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		app_idVal, ok := args["app_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: app_id"), nil
		}
		app_id, ok := app_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: app_id"), nil
		}
		release_idVal, ok := args["release_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: release_id"), nil
		}
		release_id, ok := release_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: release_id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("token=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v0.1/public/apps/%s/releases/%s/ios_manifest%s", cfg.BaseURL, app_id, release_id, queryString)
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

func CreateReleases_getiosmanifestTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v0_1_public_apps_app_id_releases_release_id_ios_manifest",
		mcp.WithDescription("Returns the manifest.plist in XML format for installing the release on a device. Only available for iOS."),
		mcp.WithString("app_id", mcp.Required(), mcp.Description("The ID of the application")),
		mcp.WithNumber("release_id", mcp.Required(), mcp.Description("The release_id")),
		mcp.WithString("token", mcp.Required(), mcp.Description("A hash that authorizes the download if it matches the release info.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Releases_getiosmanifestHandler(cfg),
	}
}

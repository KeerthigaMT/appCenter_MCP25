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

func Storereleasepublishlogs_getHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		store_nameVal, ok := args["store_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: store_name"), nil
		}
		store_name, ok := store_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: store_name"), nil
		}
		release_idVal, ok := args["release_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: release_id"), nil
		}
		release_id, ok := release_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: release_id"), nil
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
		url := fmt.Sprintf("%s/v0.1/apps/%s/%s/distribution_stores/%s/releases/%s/publish_logs", cfg.BaseURL, store_name, release_id, owner_name, app_name)
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

func CreateStorereleasepublishlogs_getTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v0_1_apps_owner_name_app_name_distribution_stores_store_name_releases_release_id_publish_logs",
		mcp.WithDescription("Returns publish logs for a particular release published to a particular store"),
		mcp.WithString("store_name", mcp.Required(), mcp.Description("The name of the store")),
		mcp.WithString("release_id", mcp.Required(), mcp.Description("The ID of the realease")),
		mcp.WithString("owner_name", mcp.Required(), mcp.Description("The name of the owner")),
		mcp.WithString("app_name", mcp.Required(), mcp.Description("The name of the application")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Storereleasepublishlogs_getHandler(cfg),
	}
}

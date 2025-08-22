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

func Releases_getlatestbypublicdistributiongroupHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		distribution_group_idVal, ok := args["distribution_group_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: distribution_group_id"), nil
		}
		distribution_group_id, ok := distribution_group_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: distribution_group_id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["is_install_page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_install_page=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v0.1/public/sdk/apps/%s/distribution_groups/%s/releases/latest%s", cfg.BaseURL, app_secret, distribution_group_id, queryString)
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

func CreateReleases_getlatestbypublicdistributiongroupTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v0_1_public_sdk_apps_app_secret_distribution_groups_distribution_group_id_releases_latest",
		mcp.WithDescription("Get a release with 'latest' for the given public group."),
		mcp.WithString("app_secret", mcp.Required(), mcp.Description("The secret of the target application")),
		mcp.WithString("distribution_group_id", mcp.Required(), mcp.Description("the id for destination")),
		mcp.WithBoolean("is_install_page", mcp.Description("The check if the request is from Install page")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Releases_getlatestbypublicdistributiongroupHandler(cfg),
	}
}

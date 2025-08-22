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

func Crashes_listHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		crash_group_idVal, ok := args["crash_group_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: crash_group_id"), nil
		}
		crash_group_id, ok := crash_group_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: crash_group_id"), nil
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
		queryParams := make([]string, 0)
		if val, ok := args["include_report"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_report=%v", val))
		}
		if val, ok := args["include_log"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_log=%v", val))
		}
		if val, ok := args["date_from"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("date_from=%v", val))
		}
		if val, ok := args["date_to"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("date_to=%v", val))
		}
		if val, ok := args["app_version"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("app_version=%v", val))
		}
		if val, ok := args["error_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("error_type=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v0.1/apps/%s/%s/crash_groups/%s/crashes%s", cfg.BaseURL, crash_group_id, owner_name, app_name, queryString)
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
		var result []Crash
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

func CreateCrashes_listTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v0_1_apps_owner_name_app_name_crash_groups_crash_group_id_crashes",
		mcp.WithDescription("Gets all crashes of a group."),
		mcp.WithString("crash_group_id", mcp.Required(), mcp.Description("id of a specific group")),
		mcp.WithBoolean("include_report", mcp.Description("true if the crash should include the raw crash report. Default is false")),
		mcp.WithBoolean("include_log", mcp.Description("true if the crash should include the custom log report. Default is false")),
		mcp.WithString("date_from", mcp.Description("")),
		mcp.WithString("date_to", mcp.Description("")),
		mcp.WithString("app_version", mcp.Description("version")),
		mcp.WithString("error_type", mcp.Description("")),
		mcp.WithString("owner_name", mcp.Required(), mcp.Description("The name of the owner")),
		mcp.WithString("app_name", mcp.Required(), mcp.Description("The name of the application")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Crashes_listHandler(cfg),
	}
}

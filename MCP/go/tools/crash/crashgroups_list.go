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

func Crashgroups_listHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
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
		if val, ok := args["last_occurrence_from"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("last_occurrence_from=%v", val))
		}
		if val, ok := args["last_occurrence_to"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("last_occurrence_to=%v", val))
		}
		if val, ok := args["app_version"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("app_version=%v", val))
		}
		if val, ok := args["group_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("group_type=%v", val))
		}
		if val, ok := args["group_status"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("group_status=%v", val))
		}
		if val, ok := args["group_text_search"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("group_text_search=%v", val))
		}
		if val, ok := args["$orderby"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("$orderby=%v", val))
		}
		if val, ok := args["continuation_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("continuation_token=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v0.1/apps/%s/%s/crash_groups%s", cfg.BaseURL, owner_name, app_name, queryString)
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

func CreateCrashgroups_listTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v0_1_apps_owner_name_app_name_crash_groups",
		mcp.WithDescription("Gets a list of crash groups and whether the list contains all available groups."),
		mcp.WithString("last_occurrence_from", mcp.Description("Earliest date when the last time a crash occured in a crash group")),
		mcp.WithString("last_occurrence_to", mcp.Description("Latest date when the last time a crash occured in a crash group")),
		mcp.WithString("app_version", mcp.Description("version")),
		mcp.WithString("group_type", mcp.Description("")),
		mcp.WithString("group_status", mcp.Description("")),
		mcp.WithString("group_text_search", mcp.Description("A freetext search that matches in crash, crash types, crash stack_traces and crash user")),
		mcp.WithString("$orderby", mcp.Description("the OData-like $orderby argument")),
		mcp.WithString("continuation_token", mcp.Description("Cassandra request continuation token. The token is used for pagination.")),
		mcp.WithString("owner_name", mcp.Required(), mcp.Description("The name of the owner")),
		mcp.WithString("app_name", mcp.Required(), mcp.Description("The name of the application")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Crashgroups_listHandler(cfg),
	}
}

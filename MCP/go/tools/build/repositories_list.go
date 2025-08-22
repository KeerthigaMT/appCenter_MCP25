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

func Repositories_listHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		source_hostVal, ok := args["source_host"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: source_host"), nil
		}
		source_host, ok := source_hostVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: source_host"), nil
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
		if val, ok := args["vstsAccountName"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("vstsAccountName=%v", val))
		}
		if val, ok := args["vstsProjectId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("vstsProjectId=%v", val))
		}
		if val, ok := args["service_connection_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("service_connection_id=%v", val))
		}
		if val, ok := args["form"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("form=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v0.1/apps/%s/%s/source_hosts/%s/repositories%s", cfg.BaseURL, source_host, owner_name, app_name, queryString)
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

func CreateRepositories_listTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v0_1_apps_owner_name_app_name_source_hosts_source_host_repositories",
		mcp.WithDescription("Gets the repositories available from the source code host"),
		mcp.WithString("source_host", mcp.Required(), mcp.Description("The source host")),
		mcp.WithString("vstsAccountName", mcp.Description("Filter repositories only for specified account and project, \"vstsProjectId\" is required")),
		mcp.WithString("vstsProjectId", mcp.Description("Filter repositories only for specified account and project, \"vstsAccountName\" is required")),
		mcp.WithString("service_connection_id", mcp.Description("The id of the service connection (private). Required for GitLab self-hosted repositories")),
		mcp.WithString("form", mcp.Description("The selected form of the object")),
		mcp.WithString("owner_name", mcp.Required(), mcp.Description("The name of the owner")),
		mcp.WithString("app_name", mcp.Required(), mcp.Description("The name of the application")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Repositories_listHandler(cfg),
	}
}

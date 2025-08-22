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

func Errors_availableversionsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["start"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("start=%v", val))
		}
		if val, ok := args["end"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("end=%v", val))
		}
		if val, ok := args["$top"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("$top=%v", val))
		}
		if val, ok := args["$skip"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("$skip=%v", val))
		}
		if val, ok := args["$filter"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("$filter=%v", val))
		}
		if val, ok := args["$inlinecount"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("$inlinecount=%v", val))
		}
		if val, ok := args["errorType"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("errorType=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v0.1/apps/%s/%s/errors/available_versions%s", cfg.BaseURL, owner_name, app_name, queryString)
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

func CreateErrors_availableversionsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v0_1_apps_owner_name_app_name_errors_available_versions",
		mcp.WithDescription("Get all available versions in the time range."),
		mcp.WithString("start", mcp.Required(), mcp.Description("Start date time in data in ISO 8601 date time format")),
		mcp.WithString("end", mcp.Description("Last date time in data in ISO 8601 date time format")),
		mcp.WithNumber("$top", mcp.Description("The maximum number of results to return. (0 will fetch all results till the max number.)")),
		mcp.WithNumber("$skip", mcp.Description("The offset (starting at 0) of the first result to return. This parameter along with limit is used to perform pagination.")),
		mcp.WithString("$filter", mcp.Description("A filter as specified in https://github.com/Microsoft/api-guidelines/blob/master/Guidelines.md#97-filtering.")),
		mcp.WithString("$inlinecount", mcp.Description("Controls whether or not to include a count of all the items across all pages.")),
		mcp.WithString("errorType", mcp.Description("Type of error (handled vs unhandled), including All")),
		mcp.WithString("owner_name", mcp.Required(), mcp.Description("The name of the owner")),
		mcp.WithString("app_name", mcp.Required(), mcp.Description("The name of the application")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Errors_availableversionsHandler(cfg),
	}
}

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

func Analytics_eventpropertycountsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		event_nameVal, ok := args["event_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: event_name"), nil
		}
		event_name, ok := event_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: event_name"), nil
		}
		event_property_nameVal, ok := args["event_property_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: event_property_name"), nil
		}
		event_property_name, ok := event_property_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: event_property_name"), nil
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
		if val, ok := args["versions"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("versions=%v", val))
		}
		if val, ok := args["$top"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("$top=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v0.1/apps/%s/%s/analytics/events/%s/properties/%s/counts%s", cfg.BaseURL, event_name, event_property_name, owner_name, app_name, queryString)
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

func CreateAnalytics_eventpropertycountsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v0_1_apps_owner_name_app_name_analytics_events_event_name_properties_event_property_name_counts",
		mcp.WithDescription("Event properties value counts during the time range in descending order."),
		mcp.WithString("event_name", mcp.Required(), mcp.Description("The id of the event.")),
		mcp.WithString("event_property_name", mcp.Required(), mcp.Description("The id of the event property.")),
		mcp.WithString("start", mcp.Required(), mcp.Description("Start date time in data in ISO 8601 date time format.")),
		mcp.WithString("end", mcp.Description("Last date time in data in ISO 8601 date time format.")),
		mcp.WithArray("versions", mcp.Description("To select specific application versions")),
		mcp.WithNumber("$top", mcp.Description("The number of property values to return. Set to 0 in order to fetch all results available.")),
		mcp.WithString("owner_name", mcp.Required(), mcp.Description("The name of the owner")),
		mcp.WithString("app_name", mcp.Required(), mcp.Description("The name of the application")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Analytics_eventpropertycountsHandler(cfg),
	}
}

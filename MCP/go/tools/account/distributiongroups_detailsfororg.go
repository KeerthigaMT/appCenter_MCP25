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

func Distributiongroups_detailsfororgHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		org_nameVal, ok := args["org_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: org_name"), nil
		}
		org_name, ok := org_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: org_name"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["apps_limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("apps_limit=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v0.1/orgs/%s/distribution_groups_details%s", cfg.BaseURL, org_name, queryString)
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

func CreateDistributiongroups_detailsfororgTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v0_1_orgs_org_name_distribution_groups_details",
		mcp.WithDescription("Returns a list of distribution groups with details for an organization"),
		mcp.WithString("org_name", mcp.Required(), mcp.Description("The organization's name")),
		mcp.WithString("apps_limit", mcp.Description("The max number of apps to include in the response")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Distributiongroups_detailsfororgHandler(cfg),
	}
}

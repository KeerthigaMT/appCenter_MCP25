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

func Billingaggregatedinformation_getallHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["service"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("service=%v", val))
		}
		if val, ok := args["period"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("period=%v", val))
		}
		if val, ok := args["showOriginalPlans"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("showOriginalPlans=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v0.1/billing/allAccountsAggregated%s", cfg.BaseURL, queryString)
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

func CreateBillingaggregatedinformation_getallTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v0_1_billing_allAccountsAggregated",
		mcp.WithDescription("Aggregated Billing Information for the requesting user and the organizations in which the user is an admin."),
		mcp.WithString("service", mcp.Description("Type of service that should be included in the Billing Information")),
		mcp.WithString("period", mcp.Description("Type of period that should be included in the Billing Information")),
		mcp.WithBoolean("showOriginalPlans", mcp.Description("Controls whether the API should show the original plan when Azure Subscription is not enabled")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Billingaggregatedinformation_getallHandler(cfg),
	}
}

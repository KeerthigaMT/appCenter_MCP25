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

func Datasubjectright_deletestatusrequestHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		tokenVal, ok := args["token"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: token"), nil
		}
		token, ok := tokenVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: token"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["email"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("email=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v0.1/user/dsr/delete/%s%s", cfg.BaseURL, token, queryString)
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

func CreateDatasubjectright_deletestatusrequestTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v0_1_user_dsr_delete_token",
		mcp.WithDescription(""),
		mcp.WithString("token", mcp.Required(), mcp.Description("Unique request ID (GUID)")),
		mcp.WithString("email", mcp.Required(), mcp.Description("Email used for delete with x-authz-bypass headers")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Datasubjectright_deletestatusrequestHandler(cfg),
	}
}

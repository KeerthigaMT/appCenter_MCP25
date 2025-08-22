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

func Legacycodepushacquisition_updatecheckHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["deploymentKey"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("deploymentKey=%v", val))
		}
		if val, ok := args["appVersion"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("appVersion=%v", val))
		}
		if val, ok := args["packageHash"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("packageHash=%v", val))
		}
		if val, ok := args["label"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("label=%v", val))
		}
		if val, ok := args["clientUniqueId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("clientUniqueId=%v", val))
		}
		if val, ok := args["isCompanion"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("isCompanion=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v0.1/legacy/updateCheck%s", cfg.BaseURL, queryString)
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

func CreateLegacycodepushacquisition_updatecheckTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v0_1_legacy_updateCheck",
		mcp.WithDescription("Check for updates"),
		mcp.WithString("deploymentKey", mcp.Description("")),
		mcp.WithString("appVersion", mcp.Description("")),
		mcp.WithString("packageHash", mcp.Description("")),
		mcp.WithString("label", mcp.Description("")),
		mcp.WithString("clientUniqueId", mcp.Description("")),
		mcp.WithString("isCompanion", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Legacycodepushacquisition_updatecheckHandler(cfg),
	}
}

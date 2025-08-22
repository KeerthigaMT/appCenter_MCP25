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

func Codepushacquisition_updatecheckHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["deployment_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("deployment_key=%v", val))
		}
		if val, ok := args["app_version"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("app_version=%v", val))
		}
		if val, ok := args["package_hash"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("package_hash=%v", val))
		}
		if val, ok := args["label"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("label=%v", val))
		}
		if val, ok := args["client_unique_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("client_unique_id=%v", val))
		}
		if val, ok := args["is_companion"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_companion=%v", val))
		}
		if val, ok := args["previous_label_or_app_version"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("previous_label_or_app_version=%v", val))
		}
		if val, ok := args["previous_deployment_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("previous_deployment_key=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v0.1/public/codepush/update_check%s", cfg.BaseURL, queryString)
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

func CreateCodepushacquisition_updatecheckTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v0_1_public_codepush_update_check",
		mcp.WithDescription("Check for updates"),
		mcp.WithString("deployment_key", mcp.Required(), mcp.Description("")),
		mcp.WithString("app_version", mcp.Required(), mcp.Description("")),
		mcp.WithString("package_hash", mcp.Description("")),
		mcp.WithString("label", mcp.Description("")),
		mcp.WithString("client_unique_id", mcp.Description("")),
		mcp.WithBoolean("is_companion", mcp.Description("")),
		mcp.WithString("previous_label_or_app_version", mcp.Description("")),
		mcp.WithString("previous_deployment_key", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Codepushacquisition_updatecheckHandler(cfg),
	}
}

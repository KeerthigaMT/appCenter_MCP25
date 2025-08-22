package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/app-center-client/mcp-server/config"
	"github.com/app-center-client/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Devices_removeuserdeviceHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		device_udidVal, ok := args["device_udid"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: device_udid"), nil
		}
		device_udid, ok := device_udidVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: device_udid"), nil
		}
		url := fmt.Sprintf("%s/v0.1/user/devices/%s", cfg.BaseURL, device_udid)
		req, err := http.NewRequest("DELETE", url, nil)
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

func CreateDevices_removeuserdeviceTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_v0_1_user_devices_device_udid",
		mcp.WithDescription("Removes an existing device from a user"),
		mcp.WithString("device_udid", mcp.Required(), mcp.Description("The UDID of the device")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Devices_removeuserdeviceHandler(cfg),
	}
}

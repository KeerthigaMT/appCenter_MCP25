package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/app-center-client/mcp-server/config"
	"github.com/app-center-client/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Devices_registeruserfordeviceHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		user_idVal, ok := args["user_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: user_id"), nil
		}
		user_id, ok := user_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: user_id"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/v0.1/users/%s/devices/register", cfg.BaseURL, user_id)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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

func CreateDevices_registeruserfordeviceTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v0_1_users_user_id_devices_register",
		mcp.WithDescription("Registers a user for an existing device"),
		mcp.WithString("user_id", mcp.Required(), mcp.Description("The ID of the user")),
		mcp.WithString("os_build", mcp.Description("Input parameter: The build number of the last known OS version running on the device")),
		mcp.WithString("os_version", mcp.Description("Input parameter: The last known OS version running on the device")),
		mcp.WithString("owner_id", mcp.Description("Input parameter: The user ID of the device owner.")),
		mcp.WithString("serial", mcp.Description("Input parameter: The device's serial number. Always empty or undefined at present.")),
		mcp.WithString("udid", mcp.Required(), mcp.Description("Input parameter: The Unique Device IDentifier of the device")),
		mcp.WithString("imei", mcp.Description("Input parameter: The device's International Mobile Equipment Identity number. Always empty or undefined at present.")),
		mcp.WithString("model", mcp.Required(), mcp.Description("Input parameter: The model identifier of the device, in the format iDeviceM,N")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Devices_registeruserfordeviceHandler(cfg),
	}
}

package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/marketcheck-apis/mcp-server/config"
	"github.com/marketcheck-apis/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func DecodevianeovinHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		vinVal, ok := args["vin"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: vin"), nil
		}
		vin, ok := vinVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: vin"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["api_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("api_key=%v", val))
		}
		if val, ok := args["include_generic"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_generic=%v", val))
		}
		if val, ok := args["force_decode"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("force_decode=%v", val))
		}
		// Handle multiple authentication parameters
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("api_key=%s", cfg.APIKey))
		}
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("append_api_key=%s", cfg.APIKey))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/decode/car/neovin/%s/specs%s", cfg.BaseURL, vin, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
		// API key already added to query string
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
		var result models.NeoVIN
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

func CreateDecodevianeovinTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_decode_car_neovin_vin_specs",
		mcp.WithDescription("NeoVIN Decoder"),
		mcp.WithString("api_key", mcp.Description("The API Authentication Key. Mandatory with all API calls.")),
		mcp.WithString("vin", mcp.Required(), mcp.Description("The VIN to identify the car. Must be a valid 17 char VIN")),
		mcp.WithBoolean("include_generic", mcp.Description("Boolean variable to indicate wheather to include generic data as well in response")),
		mcp.WithBoolean("force_decode", mcp.Description("Decode VIN on the fly instead of cached response")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DecodevianeovinHandler(cfg),
	}
}

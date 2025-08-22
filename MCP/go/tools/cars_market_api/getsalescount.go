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

func GetsalescountHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["api_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("api_key=%v", val))
		}
		if val, ok := args["car_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("car_type=%v", val))
		}
		if val, ok := args["make"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("make=%v", val))
		}
		if val, ok := args["mm"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("mm=%v", val))
		}
		if val, ok := args["ymm"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ymm=%v", val))
		}
		if val, ok := args["ymmt"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ymmt=%v", val))
		}
		if val, ok := args["taxonomy_vin"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("taxonomy_vin=%v", val))
		}
		if val, ok := args["state"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("state=%v", val))
		}
		if val, ok := args["city_state"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("city_state=%v", val))
		}
		if val, ok := args["vin"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("vin=%v", val))
		}
		if val, ok := args["country"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("country=%v", val))
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
		url := fmt.Sprintf("%s/sales/car%s", cfg.BaseURL, queryString)
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
		var result models.Sales
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

func CreateGetsalescountTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_sales_car",
		mcp.WithDescription("Get sales count by make, model, year, trim or taxonomy vin"),
		mcp.WithString("api_key", mcp.Description("The API Authentication Key. Mandatory with all API calls.")),
		mcp.WithString("car_type", mcp.Description("Inventory type for which sales count is to be searched, default is used")),
		mcp.WithString("make", mcp.Description("Make for which sales count is to be searched")),
		mcp.WithString("mm", mcp.Description("Make-Model for which sales count is to be searched, pipe seperated like mm=ford|f-150")),
		mcp.WithString("ymm", mcp.Description("Year-Make-Model for which sales count is to be searched, pipe seperated like ymm=2015|ford|f-150")),
		mcp.WithString("ymmt", mcp.Description("Year-Make-Model-Trim for which sales count is to be searched, pipe seperated like ymmt=2015|ford|f-150|platinum")),
		mcp.WithString("taxonomy_vin", mcp.Description("taxonomy_vin for which sales count is to be searched")),
		mcp.WithString("state", mcp.Description("State level sales count")),
		mcp.WithString("city_state", mcp.Description("City level sales count, pipe seperated like city_state=jacksonville|FL")),
		mcp.WithString("vin", mcp.Description("VIN that will be transformed to taxonomy_vin")),
		mcp.WithString("country", mcp.Description("Country for which the sales records are to be searched")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetsalescountHandler(cfg),
	}
}

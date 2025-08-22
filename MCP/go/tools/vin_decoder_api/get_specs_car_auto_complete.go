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

func Get_specs_car_auto_completeHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["api_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("api_key=%v", val))
		}
		if val, ok := args["field"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("field=%v", val))
		}
		if val, ok := args["input"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("input=%v", val))
		}
		if val, ok := args["year"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("year=%v", val))
		}
		if val, ok := args["make"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("make=%v", val))
		}
		if val, ok := args["model"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("model=%v", val))
		}
		if val, ok := args["trim"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("trim=%v", val))
		}
		if val, ok := args["body_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("body_type=%v", val))
		}
		if val, ok := args["body_subtype"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("body_subtype=%v", val))
		}
		if val, ok := args["vehicle_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("vehicle_type=%v", val))
		}
		if val, ok := args["transmission"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("transmission=%v", val))
		}
		if val, ok := args["drivetrain"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("drivetrain=%v", val))
		}
		if val, ok := args["fuel_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fuel_type=%v", val))
		}
		if val, ok := args["engine"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("engine=%v", val))
		}
		if val, ok := args["engine_size"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("engine_size=%v", val))
		}
		if val, ok := args["engine_block"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("engine_block=%v", val))
		}
		if val, ok := args["ignore_case"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ignore_case=%v", val))
		}
		if val, ok := args["facet_min_count"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("facet_min_count=%v", val))
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
		url := fmt.Sprintf("%s/specs/car/auto-complete%s", cfg.BaseURL, queryString)
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
		var result models.SpecsAutoCompleteResponse
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

func CreateGet_specs_car_auto_completeTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_specs_car_auto-complete",
		mcp.WithDescription("API for auto-completion of inputs based on taxonomy"),
		mcp.WithString("api_key", mcp.Description("The API Authentication Key. Mandatory with all API calls.")),
		mcp.WithString("field", mcp.Required(), mcp.Description("Field name for which you want auto-completion")),
		mcp.WithString("input", mcp.Required(), mcp.Description("Input entered so far")),
		mcp.WithString("year", mcp.Description("To filter listing on their year")),
		mcp.WithString("make", mcp.Description("To filter listings on their make")),
		mcp.WithString("model", mcp.Description("To filter listings on their model")),
		mcp.WithString("trim", mcp.Description("To filter listing on their trim")),
		mcp.WithString("body_type", mcp.Description("To filter listing on their body type")),
		mcp.WithString("body_subtype", mcp.Description("Body subtype to filter the listings on. Valid filter values are those that our Search facets API returns for unique body subtypes. You can pass in multiple body subtype values comma separated")),
		mcp.WithString("vehicle_type", mcp.Description("To filter listing on their vehicle type")),
		mcp.WithString("transmission", mcp.Description("To filter listing on their transmission")),
		mcp.WithString("drivetrain", mcp.Description("To filter listing on their drivetrain")),
		mcp.WithString("fuel_type", mcp.Description("To filter listing on their fuel type")),
		mcp.WithString("engine", mcp.Description("To filter listing on their engine")),
		mcp.WithString("engine_size", mcp.Description("Engine Size to match. Valid filter values are those that our Search facets API returns for unique engine size. You can pass in multiple engine size values comma separated")),
		mcp.WithString("engine_block", mcp.Description("Engine Block to match. Valid filter values are those that our Search facets API returns for unique Engine Block. You can pass in multiple Engine Block values comma separated")),
		mcp.WithBoolean("ignore_case", mcp.Description("Boolean variable to indicate ignore case of current input")),
		mcp.WithString("facet_min_count", mcp.Description("Provide minimum count value for facets")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_specs_car_auto_completeHandler(cfg),
	}
}

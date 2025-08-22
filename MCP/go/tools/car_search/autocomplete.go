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

func AutocompleteHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["exterior_color"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("exterior_color=%v", val))
		}
		if val, ok := args["interior_color"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("interior_color=%v", val))
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
		if val, ok := args["state"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("state=%v", val))
		}
		if val, ok := args["city"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("city=%v", val))
		}
		if val, ok := args["source"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("source=%v", val))
		}
		if val, ok := args["dealer_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dealer_id=%v", val))
		}
		if val, ok := args["country"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("country=%v", val))
		}
		if val, ok := args["car_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("car_type=%v", val))
		}
		if val, ok := args["include_non_vin_listings"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_non_vin_listings=%v", val))
		}
		if val, ok := args["ignore_case"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ignore_case=%v", val))
		}
		if val, ok := args["term_counts"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("term_counts=%v", val))
		}
		if val, ok := args["sort_by"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_by=%v", val))
		}
		if val, ok := args["seller_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("seller_type=%v", val))
		}
		if val, ok := args["radius"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("radius=%v", val))
		}
		if val, ok := args["zip"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("zip=%v", val))
		}
		if val, ok := args["inventory_count_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("inventory_count_range=%v", val))
		}
		if val, ok := args["exclude_dealer_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("exclude_dealer_ids=%v", val))
		}
		if val, ok := args["exclude_sources"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("exclude_sources=%v", val))
		}
		if val, ok := args["in_transit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("in_transit=%v", val))
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
		url := fmt.Sprintf("%s/search/car/auto-complete%s", cfg.BaseURL, queryString)
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
		var result models.SearchAutoCompleteResponse
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

func CreateAutocompleteTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_search_car_auto-complete",
		mcp.WithDescription("API for auto-completion of inputs"),
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
		mcp.WithString("exterior_color", mcp.Description("Exterior color to match. Valid filter values are those that our Search facets API returns for unique exterior colors. You can pass in multiple exterior color values comma separated")),
		mcp.WithString("interior_color", mcp.Description("Interior color to match. Valid filter values are those that our Search facets API returns for unique interior colors. You can pass in multiple interior color values comma separated")),
		mcp.WithString("engine", mcp.Description("To filter listing on their engine")),
		mcp.WithString("engine_size", mcp.Description("Engine Size to match. Valid filter values are those that our Search facets API returns for unique engine size. You can pass in multiple engine size values comma separated")),
		mcp.WithString("engine_block", mcp.Description("Engine Block to match. Valid filter values are those that our Search facets API returns for unique Engine Block. You can pass in multiple Engine Block values comma separated")),
		mcp.WithString("state", mcp.Description("To filter listing on State in which they are listed")),
		mcp.WithString("city", mcp.Description("To filter listing on City in which they are listed")),
		mcp.WithString("source", mcp.Description("To filter listing on their source only for widget requests")),
		mcp.WithString("dealer_id", mcp.Description("Dealer id to filter the listings.")),
		mcp.WithString("country", mcp.Description("To filter listing on Country in which they are listed")),
		mcp.WithString("car_type", mcp.Description("Car type. Allowed values are - new / used")),
		mcp.WithString("include_non_vin_listings", mcp.Description("Flag to indicate whether to include non vin listing terms in results or not. Default is false to avoid un-normalised terms from non vin listings out of results")),
		mcp.WithBoolean("ignore_case", mcp.Description("Boolean variable to indicate ignore case of current input")),
		mcp.WithBoolean("term_counts", mcp.Description("Boolean variable to indicate wheather to include term counts as well in response")),
		mcp.WithString("sort_by", mcp.Description("Sort the response, either by index or count(default)")),
		mcp.WithString("seller_type", mcp.Description("seller type for autocomplete")),
		mcp.WithNumber("radius", mcp.Description("Radius around the search location (Unit - Miles)")),
		mcp.WithString("zip", mcp.Description("To filter listing on ZIP around which they are listed")),
		mcp.WithString("inventory_count_range", mcp.Description("Inventory count range to filter listings with count of total listings in dealers inventory. Range to be given in the format - min-max e.g. 10-50")),
		mcp.WithString("exclude_dealer_ids", mcp.Description("A list of dealer ids to exclude from result")),
		mcp.WithString("exclude_sources", mcp.Description("A list of sources to exclude from result")),
		mcp.WithString("in_transit", mcp.Description("A boolean to filter in transit vehicles")),
		mcp.WithString("facet_min_count", mcp.Description("Provide minimum count value for facets")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    AutocompleteHandler(cfg),
	}
}

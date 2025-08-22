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

func Get_search_rv_activeHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["api_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("api_key=%v", val))
		}
		if val, ok := args["price_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("price_range=%v", val))
		}
		if val, ok := args["miles_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("miles_range=%v", val))
		}
		if val, ok := args["msrp_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("msrp_range=%v", val))
		}
		if val, ok := args["year_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("year_range=%v", val))
		}
		if val, ok := args["search_text"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("search_text=%v", val))
		}
		if val, ok := args["latitude"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("latitude=%v", val))
		}
		if val, ok := args["longitude"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("longitude=%v", val))
		}
		if val, ok := args["radius"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("radius=%v", val))
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
		if val, ok := args["model_o"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("model_o=%v", val))
		}
		if val, ok := args["vin"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("vin=%v", val))
		}
		if val, ok := args["inventory_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("inventory_type=%v", val))
		}
		if val, ok := args["stock_no"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("stock_no=%v", val))
		}
		if val, ok := args["source"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("source=%v", val))
		}
		if val, ok := args["dealer_name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dealer_name=%v", val))
		}
		if val, ok := args["dealer_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dealer_id=%v", val))
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
		if val, ok := args["fuel_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fuel_type=%v", val))
		}
		if val, ok := args["transmission"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("transmission=%v", val))
		}
		if val, ok := args["class"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("class=%v", val))
		}
		if val, ok := args["state"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("state=%v", val))
		}
		if val, ok := args["city"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("city=%v", val))
		}
		if val, ok := args["zip"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("zip=%v", val))
		}
		if val, ok := args["msa_code"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("msa_code=%v", val))
		}
		if val, ok := args["sort_by"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_by=%v", val))
		}
		if val, ok := args["sort_order"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_order=%v", val))
		}
		if val, ok := args["rows"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("rows=%v", val))
		}
		if val, ok := args["start"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("start=%v", val))
		}
		if val, ok := args["facets"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("facets=%v", val))
		}
		if val, ok := args["range_facets"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("range_facets=%v", val))
		}
		if val, ok := args["facet_sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("facet_sort=%v", val))
		}
		if val, ok := args["stats"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("stats=%v", val))
		}
		if val, ok := args["last_seen_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("last_seen_range=%v", val))
		}
		if val, ok := args["first_seen_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("first_seen_range=%v", val))
		}
		if val, ok := args["last_seen_days"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("last_seen_days=%v", val))
		}
		if val, ok := args["first_seen_days"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("first_seen_days=%v", val))
		}
		if val, ok := args["slideouts"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("slideouts=%v", val))
		}
		if val, ok := args["length_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("length_range=%v", val))
		}
		if val, ok := args["length"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("length=%v", val))
		}
		if val, ok := args["base_exterior_color"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("base_exterior_color=%v", val))
		}
		if val, ok := args["base_interior_color"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("base_interior_color=%v", val))
		}
		if val, ok := args["seating_capacity"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("seating_capacity=%v", val))
		}
		if val, ok := args["fresh_water_capacity"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fresh_water_capacity=%v", val))
		}
		if val, ok := args["sleeps"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sleeps=%v", val))
		}
		if val, ok := args["cylinders"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cylinders=%v", val))
		}
		if val, ok := args["number_of_awnings"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("number_of_awnings=%v", val))
		}
		if val, ok := args["doors"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("doors=%v", val))
		}
		if val, ok := args["gvwr"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("gvwr=%v", val))
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
		url := fmt.Sprintf("%s/search/rv/active%s", cfg.BaseURL, queryString)
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
		var result models.RVSearchResponse
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

func CreateGet_search_rv_activeTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_search_rv_active",
		mcp.WithDescription("Gets active RV listings for the given search criteria"),
		mcp.WithString("api_key", mcp.Description("The API Authentication Key. Mandatory with all API calls.")),
		mcp.WithString("price_range", mcp.Description("Price range to filter listings with the price in the range given. Range to be given in the format - min-max e.g. 1000-5000")),
		mcp.WithString("miles_range", mcp.Description("Miles range to filter listings with miles in the given range. Range to be given in the format - min-max e.g. 1000-5000")),
		mcp.WithString("msrp_range", mcp.Description("MSRP range to filter listings with the msrp in the range given. Range to be given in the format - min-max e.g. 1000-5000")),
		mcp.WithString("year_range", mcp.Description("Year range to filter listings with the year in the range given. Range to be given in the format - min-max e.g. 2019-2021")),
		mcp.WithString("search_text", mcp.Description("To search a substring across entire document")),
		mcp.WithString("latitude", mcp.Description("Latitude component of location")),
		mcp.WithString("longitude", mcp.Description("Longitude component of location")),
		mcp.WithNumber("radius", mcp.Description("Radius around the search location (Unit - Miles)")),
		mcp.WithString("year", mcp.Description("To filter listing on their year")),
		mcp.WithString("make", mcp.Description("To filter listings on their make")),
		mcp.WithString("model", mcp.Description("To filter listings on their model")),
		mcp.WithString("model_o", mcp.Description("To filter listings on their model orig (as described on the webpage)")),
		mcp.WithString("vin", mcp.Description("To filter listing on their VIN")),
		mcp.WithString("inventory_type", mcp.Description("To filter listing on their condition. Either used or new")),
		mcp.WithString("stock_no", mcp.Description("To filter listing on their stock number on lot")),
		mcp.WithString("source", mcp.Description("To filter listing on their source")),
		mcp.WithString("dealer_name", mcp.Description("Filter listings on dealer_name")),
		mcp.WithString("dealer_id", mcp.Description("Dealer id to filter the listings.")),
		mcp.WithString("exterior_color", mcp.Description("Exterior color to match. Valid filter values are those that our Search facets API returns for unique exterior colors. You can pass in multiple exterior color values comma separated")),
		mcp.WithString("interior_color", mcp.Description("Interior color to match. Valid filter values are those that our Search facets API returns for unique interior colors. You can pass in multiple interior color values comma separated")),
		mcp.WithString("engine", mcp.Description("To filter listing on their engine")),
		mcp.WithString("fuel_type", mcp.Description("To filter listing on their fuel type")),
		mcp.WithString("transmission", mcp.Description("To filter listing on their transmission")),
		mcp.WithString("class", mcp.Description("Filter RV listings on class")),
		mcp.WithString("state", mcp.Description("To filter listing on State in which they are listed")),
		mcp.WithString("city", mcp.Description("To filter listing on City in which they are listed")),
		mcp.WithString("zip", mcp.Description("To filter listing on ZIP around which they are listed")),
		mcp.WithString("msa_code", mcp.Description("To filter listing on msa code in which they are listed")),
		mcp.WithString("sort_by", mcp.Description("Sort by field. Default sort field is distance from the given point")),
		mcp.WithString("sort_order", mcp.Description("Sort order - asc or desc. Default sort order is asc")),
		mcp.WithNumber("rows", mcp.Description("Number of results to return. Default is 10. Max is 50")),
		mcp.WithNumber("start", mcp.Description("Page number to fetch the results for the given criteria. Default is 0. Pagination is allowed only till first 10000 results for the search and sort criteria. The page value can be only between 1 to 10000/rows")),
		mcp.WithString("facets", mcp.Description("The comma separated list of fields for which facets are requested. Facets could be requested in addition to the listings for the search. Please note - The API calls with lots of facet fields may take longer to respond.")),
		mcp.WithString("range_facets", mcp.Description("The comma separated list of numeric fields for which range facets are requested. Range facets could be requested in addition to the listings for the search. Please note - The API calls with lots of range facet fields may take longer to respond.")),
		mcp.WithString("facet_sort", mcp.Description("Control sort order of facets with this parameter with default sort being on count, Other available sort is alphabetical sort, which can be obtained by using index as value for this param")),
		mcp.WithString("stats", mcp.Description("The list of fields for which stats need to be generated based on the matching listings for the search criteria. The stats consists of mean, max, average and count of listings based on which the stats are calculated for the field. Stats could be requested in addition to the listings for the search. Please note - The API calls with the stats fields may take longer to respond.")),
		mcp.WithString("last_seen_range", mcp.Description("Last seen date range to filter listings with the last seen in the range given. Range to be given in the format [YYYYMMDD] - min-max e.g. 20190523-20190623")),
		mcp.WithString("first_seen_range", mcp.Description("First seen date range to filter listings with the first seen in the range given. Range to be given in the format [YYYYMMDD] - min-max e.g. 20190523-20190623")),
		mcp.WithString("last_seen_days", mcp.Description("Last seen days range to filter listings with the last seen in the range given. Range to be given in the format - min-max e.g. 25-12")),
		mcp.WithString("first_seen_days", mcp.Description("First seen days range to filter listings with the first seen in the range given. Range to be given in the format - min-max e.g. 25-12")),
		mcp.WithString("slideouts", mcp.Description("Filter RV listings on slideouts")),
		mcp.WithString("length_range", mcp.Description("length range to filter listings with the length in the range given. Range to be given in the format - min-max e.g. 50-200")),
		mcp.WithString("length", mcp.Description("Filter RV listings on length")),
		mcp.WithString("base_exterior_color", mcp.Description("Base exterior color to match. Valid filter values are those that our Search facets API returns for unique base exterior colors. You can pass in multiple base interior color values comma separated")),
		mcp.WithString("base_interior_color", mcp.Description("Base interior color to match. Valid filter values are those that our Search facets API returns for unique base interior colors. You can pass in multiple base interior color values comma separated")),
		mcp.WithString("seating_capacity", mcp.Description("To filter on vehicle seating capacity")),
		mcp.WithString("fresh_water_capacity", mcp.Description("To filter on fresh water capacity of vehicle")),
		mcp.WithString("sleeps", mcp.Description("To filter data based on sleeps")),
		mcp.WithString("cylinders", mcp.Description("To filter listing on their cylinders")),
		mcp.WithString("number_of_awnings", mcp.Description("To filter on number_of_awnings")),
		mcp.WithString("doors", mcp.Description("Doors to filter the cars on. Valid filter values are those that our Search facets API returns for unique doors. You can pass in multiple doors values comma separated")),
		mcp.WithString("gvwr", mcp.Description("To filter on the maximum total weight of your vehicle")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_search_rv_activeHandler(cfg),
	}
}

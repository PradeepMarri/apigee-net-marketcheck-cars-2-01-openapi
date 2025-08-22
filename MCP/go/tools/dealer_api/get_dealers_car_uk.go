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

func Get_dealers_car_ukHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["api_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("api_key=%v", val))
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
		if val, ok := args["rows"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("rows=%v", val))
		}
		if val, ok := args["start"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("start=%v", val))
		}
		if val, ok := args["country"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("country=%v", val))
		}
		if val, ok := args["dealer_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dealer_type=%v", val))
		}
		if val, ok := args["city"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("city=%v", val))
		}
		if val, ok := args["county"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("county=%v", val))
		}
		if val, ok := args["listing_count_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("listing_count_range=%v", val))
		}
		if val, ok := args["inventory_url"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("inventory_url=%v", val))
		}
		if val, ok := args["postal_code"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("postal_code=%v", val))
		}
		if val, ok := args["sort_by"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_by=%v", val))
		}
		if val, ok := args["sort_order"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_order=%v", val))
		}
		if val, ok := args["provider"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("provider=%v", val))
		}
		if val, ok := args["facets"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("facets=%v", val))
		}
		if val, ok := args["range_facets"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("range_facets=%v", val))
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
		url := fmt.Sprintf("%s/dealers/car/uk%s", cfg.BaseURL, queryString)
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
		var result models.DealersResponse
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

func CreateGet_dealers_car_ukTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_dealers_car_uk",
		mcp.WithDescription("Find car dealers around"),
		mcp.WithString("api_key", mcp.Description("The API Authentication Key. Mandatory with all API calls.")),
		mcp.WithString("latitude", mcp.Description("Latitude component of location")),
		mcp.WithString("longitude", mcp.Description("Longitude component of location")),
		mcp.WithNumber("radius", mcp.Description("Radius around the search location (Unit - Miles)")),
		mcp.WithNumber("rows", mcp.Description("Number of results to return. Default is 10. Max is 50")),
		mcp.WithNumber("start", mcp.Description("Page number to fetch the results for the given criteria. Default is 0. Pagination is allowed only till first 10000 results for the search and sort criteria. The page value can be only between 1 to 10000/rows")),
		mcp.WithString("country", mcp.Description("To filter listing on Country in which they are listed")),
		mcp.WithString("dealer_type", mcp.Description("Filter based on dealer type independant or franchise")),
		mcp.WithString("city", mcp.Description("To filter listing on City in which they are listed")),
		mcp.WithString("county", mcp.Description("To filter listing on county in which they are listed")),
		mcp.WithString("listing_count_range", mcp.Description("To filter dealers based on their inventory size. Range can be given in the format - min-max e.g. 50-100")),
		mcp.WithString("inventory_url", mcp.Description("inventory_url of dealer to be searched")),
		mcp.WithString("postal_code", mcp.Description("To filter listing on postal code around which they are listed")),
		mcp.WithString("sort_by", mcp.Description("Sort by field. Default sort field is distance from the given point")),
		mcp.WithString("sort_order", mcp.Description("Sort order - asc or desc. Default sort order is asc")),
		mcp.WithBoolean("provider", mcp.Description("boolean param to include site providers name in response")),
		mcp.WithString("facets", mcp.Description("The comma separated list of fields for which facets are requested. Facets could be requested in addition to the listings for the search. Please note - The API calls with lots of facet fields may take longer to respond.")),
		mcp.WithString("range_facets", mcp.Description("The comma separated list of numeric fields for which range facets are requested. Range facets could be requested in addition to the listings for the search. Please note - The API calls with lots of range facet fields may take longer to respond.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_dealers_car_ukHandler(cfg),
	}
}

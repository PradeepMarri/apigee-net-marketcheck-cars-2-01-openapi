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

func OemsearchHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["api_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("api_key=%v", val))
		}
		if val, ok := args["offer_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("offer_type=%v", val))
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
		if val, ok := args["msrp"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("msrp=%v", val))
		}
		if val, ok := args["apr"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("apr=%v", val))
		}
		if val, ok := args["monthly"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("monthly=%v", val))
		}
		if val, ok := args["monthly_per_thousand"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("monthly_per_thousand=%v", val))
		}
		if val, ok := args["down_payment"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("down_payment=%v", val))
		}
		if val, ok := args["due_at_signing"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("due_at_signing=%v", val))
		}
		if val, ok := args["security_deposit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("security_deposit=%v", val))
		}
		if val, ok := args["disposition_fee"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("disposition_fee=%v", val))
		}
		if val, ok := args["acquisition_fee"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("acquisition_fee=%v", val))
		}
		if val, ok := args["duration"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("duration=%v", val))
		}
		if val, ok := args["dealer_contribution"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dealer_contribution=%v", val))
		}
		if val, ok := args["mileage_charge"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("mileage_charge=%v", val))
		}
		if val, ok := args["mileage_charge_limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("mileage_charge_limit=%v", val))
		}
		if val, ok := args["cashback_amount"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cashback_amount=%v", val))
		}
		if val, ok := args["cashback_target_group"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cashback_target_group=%v", val))
		}
		if val, ok := args["lease_end_purchase_option"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lease_end_purchase_option=%v", val))
		}
		if val, ok := args["net_capitalised_cost"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("net_capitalised_cost=%v", val))
		}
		if val, ok := args["gross_capitalised_cost"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("gross_capitalised_cost=%v", val))
		}
		if val, ok := args["total_monthly_payment"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("total_monthly_payment=%v", val))
		}
		if val, ok := args["zip"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("zip=%v", val))
		}
		if val, ok := args["city"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("city=%v", val))
		}
		if val, ok := args["state"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("state=%v", val))
		}
		if val, ok := args["country"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("country=%v", val))
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
		if val, ok := args["search_text"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("search_text=%v", val))
		}
		if val, ok := args["last_seen_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("last_seen_range=%v", val))
		}
		if val, ok := args["first_seen_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("first_seen_range=%v", val))
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
		url := fmt.Sprintf("%s/search/car/incentive/oem%s", cfg.BaseURL, queryString)
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
		var result models.SearchResponse
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

func CreateOemsearchTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_search_car_incentive_oem",
		mcp.WithDescription("Gets oem incentive listings for the given search criteria"),
		mcp.WithString("api_key", mcp.Description("The API Authentication Key. Mandatory with all API calls.")),
		mcp.WithString("offer_type", mcp.Description("The type of the incentive")),
		mcp.WithString("year", mcp.Description("To filter listing on their year")),
		mcp.WithString("make", mcp.Description("To filter listings on their make")),
		mcp.WithString("model", mcp.Description("To filter listings on their model")),
		mcp.WithString("trim", mcp.Description("To filter listing on their trim")),
		mcp.WithString("msrp", mcp.Description("MSRP range to filter listings with the msrp in the range given. Range to be given in the format - min-max e.g. 1000-5000")),
		mcp.WithString("apr", mcp.Description("APR range to filter listings with the msrp in the range given. Range to be given in the format - min-max e.g. 1000-5000")),
		mcp.WithString("monthly", mcp.Description("To filter listing on Monthly payment amount, usually populated in Lease offers")),
		mcp.WithString("monthly_per_thousand", mcp.Description("To filter listing on monthly amount to be paid by customer for every $1000 financed at the advertised APR rate")),
		mcp.WithString("down_payment", mcp.Description("To filter listing on down payment offer on car")),
		mcp.WithString("due_at_signing", mcp.Description("To filter listing on total amount due at signing, that usually includes first month payment, down payment, acquisition fee etc")),
		mcp.WithString("security_deposit", mcp.Description("To filter listing on security deposit required for the offer")),
		mcp.WithString("disposition_fee", mcp.Description("To filter listing on disposition fee of the car")),
		mcp.WithString("acquisition_fee", mcp.Description("To filter listing on acquisition fee of the car")),
		mcp.WithString("duration", mcp.Description("To filter listing on offer duration in months")),
		mcp.WithString("dealer_contribution", mcp.Description("To filter listing on any contribution from dealer's side")),
		mcp.WithString("mileage_charge", mcp.Description("Mileage Charge Range range to filter listings with the msrp in the range given. Range to be given in the format - min-max e.g. 100-1000")),
		mcp.WithString("mileage_charge_limit", mcp.Description("To filter listing on mileage charge limit the offer is valid up to under the default clauses")),
		mcp.WithString("cashback_amount", mcp.Description("To filter listing on cashback amounts listed in offer")),
		mcp.WithString("cashback_target_group", mcp.Description("To filter listing on the demographic or any other entity for whom this cashback offer is for. Not all target groups are identified but the most common ones are tagged like Military, Grad students Current owners etc")),
		mcp.WithString("lease_end_purchase_option", mcp.Description("To filter listing on amount at the lease end to pay for buying the car")),
		mcp.WithString("net_capitalised_cost", mcp.Description("To filter listing on net capitalised cost of the car")),
		mcp.WithString("gross_capitalised_cost", mcp.Description("To filter listing on gross capitalised cost of the car")),
		mcp.WithString("total_monthly_payment", mcp.Description("To filter listing on gross capitalised cost of the car")),
		mcp.WithString("zip", mcp.Description("To filter listing on ZIP around which they are listed")),
		mcp.WithString("city", mcp.Description("To filter listing on City in which they are listed")),
		mcp.WithString("state", mcp.Description("To filter listing on State in which they are listed")),
		mcp.WithString("country", mcp.Description("To filter listing on Country in which they are listed")),
		mcp.WithString("latitude", mcp.Description("Latitude component of location")),
		mcp.WithString("longitude", mcp.Description("Longitude component of location")),
		mcp.WithNumber("radius", mcp.Description("Radius around the search location (Unit - Miles)")),
		mcp.WithString("search_text", mcp.Description("To search a substring across entire document")),
		mcp.WithString("last_seen_range", mcp.Description("Last seen date range to filter listings with the last seen in the range given. Range to be given in the format [YYYYMMDD] - min-max e.g. 20190523-20190623")),
		mcp.WithString("first_seen_range", mcp.Description("First seen date range to filter listings with the first seen in the range given. Range to be given in the format [YYYYMMDD] - min-max e.g. 20190523-20190623")),
		mcp.WithString("sort_by", mcp.Description("Sort by field. Default sort field is distance from the given point")),
		mcp.WithString("sort_order", mcp.Description("Sort order - asc or desc. Default sort order is asc")),
		mcp.WithNumber("rows", mcp.Description("Number of results to return. Default is 10. Max is 50")),
		mcp.WithNumber("start", mcp.Description("Page number to fetch the results for the given criteria. Default is 0. Pagination is allowed only till first 10000 results for the search and sort criteria. The page value can be only between 1 to 10000/rows")),
		mcp.WithString("facets", mcp.Description("The comma separated list of fields for which facets are requested. Facets could be requested in addition to the listings for the search. Please note - The API calls with lots of facet fields may take longer to respond.")),
		mcp.WithString("range_facets", mcp.Description("The comma separated list of numeric fields for which range facets are requested. Range facets could be requested in addition to the listings for the search. Please note - The API calls with lots of range facet fields may take longer to respond.")),
		mcp.WithString("facet_sort", mcp.Description("Control sort order of facets with this parameter with default sort being on count, Other available sort is alphabetical sort, which can be obtained by using index as value for this param")),
		mcp.WithString("stats", mcp.Description("The list of fields for which stats need to be generated based on the matching listings for the search criteria. The stats consists of mean, max, average and count of listings based on which the stats are calculated for the field. Stats could be requested in addition to the listings for the search. Please note - The API calls with the stats fields may take longer to respond.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    OemsearchHandler(cfg),
	}
}

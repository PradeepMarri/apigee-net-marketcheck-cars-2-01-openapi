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

func Get_search_car_uk_recentsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["api_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("api_key=%v", val))
		}
		if val, ok := args["append_api_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("append_api_key=%v", val))
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
		if val, ok := args["zip"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("zip=%v", val))
		}
		if val, ok := args["include_lease"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_lease=%v", val))
		}
		if val, ok := args["include_finance"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_finance=%v", val))
		}
		if val, ok := args["lease_term"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lease_term=%v", val))
		}
		if val, ok := args["lease_down_payment"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lease_down_payment=%v", val))
		}
		if val, ok := args["lease_emp"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lease_emp=%v", val))
		}
		if val, ok := args["finance_loan_term"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("finance_loan_term=%v", val))
		}
		if val, ok := args["finance_loan_apr"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("finance_loan_apr=%v", val))
		}
		if val, ok := args["finance_emp"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("finance_emp=%v", val))
		}
		if val, ok := args["finance_down_payment"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("finance_down_payment=%v", val))
		}
		if val, ok := args["finance_down_payment_per"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("finance_down_payment_per=%v", val))
		}
		if val, ok := args["car_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("car_type=%v", val))
		}
		if val, ok := args["carfax_1_owner"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("carfax_1_owner=%v", val))
		}
		if val, ok := args["carfax_clean_title"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("carfax_clean_title=%v", val))
		}
		if val, ok := args["year_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("year_range=%v", val))
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
		if val, ok := args["dealer_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dealer_id=%v", val))
		}
		if val, ok := args["source"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("source=%v", val))
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
		if val, ok := args["ymmt"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ymmt=%v", val))
		}
		if val, ok := args["match"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("match=%v", val))
		}
		if val, ok := args["cylinders"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cylinders=%v", val))
		}
		if val, ok := args["transmission"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("transmission=%v", val))
		}
		if val, ok := args["doors"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("doors=%v", val))
		}
		if val, ok := args["drivetrain"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("drivetrain=%v", val))
		}
		if val, ok := args["exterior_color"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("exterior_color=%v", val))
		}
		if val, ok := args["interior_color"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("interior_color=%v", val))
		}
		if val, ok := args["base_exterior_color"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("base_exterior_color=%v", val))
		}
		if val, ok := args["base_interior_color"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("base_interior_color=%v", val))
		}
		if val, ok := args["engine"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("engine=%v", val))
		}
		if val, ok := args["engine_size"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("engine_size=%v", val))
		}
		if val, ok := args["engine_aspiration"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("engine_aspiration=%v", val))
		}
		if val, ok := args["engine_block"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("engine_block=%v", val))
		}
		if val, ok := args["highway_mpg_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("highway_mpg_range=%v", val))
		}
		if val, ok := args["city_mpg_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("city_mpg_range=%v", val))
		}
		if val, ok := args["combined_mpg_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("combined_mpg_range=%v", val))
		}
		if val, ok := args["miles_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("miles_range=%v", val))
		}
		if val, ok := args["price_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("price_range=%v", val))
		}
		if val, ok := args["msrp_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("msrp_range=%v", val))
		}
		if val, ok := args["dom_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dom_range=%v", val))
		}
		if val, ok := args["last_seen_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("last_seen_range=%v", val))
		}
		if val, ok := args["first_seen_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("first_seen_range=%v", val))
		}
		if val, ok := args["first_seen_at_source_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("first_seen_at_source_range=%v", val))
		}
		if val, ok := args["first_seen_at_mc_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("first_seen_at_mc_range=%v", val))
		}
		if val, ok := args["last_seen_days"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("last_seen_days=%v", val))
		}
		if val, ok := args["first_seen_days"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("first_seen_days=%v", val))
		}
		if val, ok := args["first_seen_at_source_days"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("first_seen_at_source_days=%v", val))
		}
		if val, ok := args["first_seen_at_mc_days"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("first_seen_at_mc_days=%v", val))
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
		if val, ok := args["include_non_vin_listings"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_non_vin_listings=%v", val))
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
		if val, ok := args["country"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("country=%v", val))
		}
		if val, ok := args["plot"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("plot=%v", val))
		}
		if val, ok := args["nodedup"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("nodedup=%v", val))
		}
		if val, ok := args["dedup"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dedup=%v", val))
		}
		if val, ok := args["owned"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("owned=%v", val))
		}
		if val, ok := args["state"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("state=%v", val))
		}
		if val, ok := args["city"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("city=%v", val))
		}
		if val, ok := args["msa_code"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("msa_code=%v", val))
		}
		if val, ok := args["dealer_name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dealer_name=%v", val))
		}
		if val, ok := args["dealership_group_name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dealership_group_name=%v", val))
		}
		if val, ok := args["trim_o"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("trim_o=%v", val))
		}
		if val, ok := args["trim_r"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("trim_r=%v", val))
		}
		if val, ok := args["dom_active_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dom_active_range=%v", val))
		}
		if val, ok := args["dom_180_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dom_180_range=%v", val))
		}
		if val, ok := args["exclude_certified"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("exclude_certified=%v", val))
		}
		if val, ok := args["fuel_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fuel_type=%v", val))
		}
		if val, ok := args["dealer_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dealer_type=%v", val))
		}
		if val, ok := args["photo_links"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("photo_links=%v", val))
		}
		if val, ok := args["photo_links_cached"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("photo_links_cached=%v", val))
		}
		if val, ok := args["stock_no"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("stock_no=%v", val))
		}
		if val, ok := args["sold"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sold=%v", val))
		}
		if val, ok := args["include_relevant_links"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_relevant_links=%v", val))
		}
		if val, ok := args["expired"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("expired=%v", val))
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
		if val, ok := args["seating_capacity"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("seating_capacity=%v", val))
		}
		if val, ok := args["insurance_group"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("insurance_group=%v", val))
		}
		if val, ok := args["vrm"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("vrm=%v", val))
		}
		if val, ok := args["num_owners"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("num_owners=%v", val))
		}
		if val, ok := args["variant"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("variant=%v", val))
		}
		if val, ok := args["postal_code"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("postal_code=%v", val))
		}
		if val, ok := args["write_off_category"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("write_off_category=%v", val))
		}
		if val, ok := args["fca_status"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fca_status=%v", val))
		}
		if val, ok := args["active_inventory_date_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("active_inventory_date_range=%v", val))
		}
		if val, ok := args["engine_size_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("engine_size_range=%v", val))
		}
		if val, ok := args["price_change_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("price_change_range=%v", val))
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
		url := fmt.Sprintf("%s/search/car/uk/recents%s", cfg.BaseURL, queryString)
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
		var result models.UKSearchResponse
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

func CreateGet_search_car_uk_recentsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_search_car_uk_recents",
		mcp.WithDescription("Gets Recent UK car listings for the given search criteria"),
		mcp.WithString("api_key", mcp.Description("The API Authentication Key. Mandatory with all API calls.")),
		mcp.WithBoolean("append_api_key", mcp.Description("Flag on whether to include api_key in response API urls (if any)")),
		mcp.WithString("latitude", mcp.Description("Latitude component of location")),
		mcp.WithString("longitude", mcp.Description("Longitude component of location")),
		mcp.WithNumber("radius", mcp.Description("Radius around the search location (Unit - Miles)")),
		mcp.WithString("zip", mcp.Description("To filter listing on ZIP around which they are listed")),
		mcp.WithBoolean("include_lease", mcp.Description("Boolean param to search for listings that include leasing options in them")),
		mcp.WithBoolean("include_finance", mcp.Description("Boolean param to search for listings that include finance options in them")),
		mcp.WithString("lease_term", mcp.Description("Search listings with exact lease term, or inside a range with min and max seperated by a dash like lease_term=30-60")),
		mcp.WithString("lease_down_payment", mcp.Description("Search listings with exact down payment in lease offers, or inside a range with min and max seperated by a dash like lease_down_payment=30-60")),
		mcp.WithString("lease_emp", mcp.Description("Search listings with lease offers exactly matching Estimated Monthly Payment(EMI), or inside a range with min and max seperated by a dash like lease_emp=30-60")),
		mcp.WithString("finance_loan_term", mcp.Description("Search listings with exact finance loan term, or inside a range with min and max seperated by a dash like finance_loan_term=30-60")),
		mcp.WithString("finance_loan_apr", mcp.Description("Search listings with finance offers exactly matching loans Annual Percentage Rate, or inside a range with min and max seperated by a dash like finance_loan_apr=30-60")),
		mcp.WithString("finance_emp", mcp.Description("Search listings with finance offers exactly matching Estimated Monthly Payment(EMI), or inside a range with min and max seperated by a dash like finance_emp=30-60")),
		mcp.WithString("finance_down_payment", mcp.Description("Search listings with exact down payment in finance offers, or inside a range with min and max seperated by a dash like finance_down_payment=30-60")),
		mcp.WithString("finance_down_payment_per", mcp.Description("Search listings with exact down payment percentage in finance offers, or inside a range with min and max seperated by a dash like finance_down_payment_per=30-60")),
		mcp.WithString("car_type", mcp.Description("Car type. Allowed values are - new / used / certified")),
		mcp.WithString("carfax_1_owner", mcp.Description("Indicates whether car has had only one owner or not")),
		mcp.WithString("carfax_clean_title", mcp.Description("Indicates whether car has clean ownership records")),
		mcp.WithString("year_range", mcp.Description("Year range to filter listings with the year in the range given. Range to be given in the format - min-max e.g. 2019-2021")),
		mcp.WithString("year", mcp.Description("To filter listing on their year")),
		mcp.WithString("make", mcp.Description("To filter listings on their make")),
		mcp.WithString("model", mcp.Description("To filter listings on their model")),
		mcp.WithString("trim", mcp.Description("To filter listing on their trim")),
		mcp.WithString("dealer_id", mcp.Description("Dealer id to filter the listings.")),
		mcp.WithString("source", mcp.Description("To filter listing on their source")),
		mcp.WithString("body_type", mcp.Description("To filter listing on their body type")),
		mcp.WithString("body_subtype", mcp.Description("Body subtype to filter the listings on. Valid filter values are those that our Search facets API returns for unique body subtypes. You can pass in multiple body subtype values comma separated")),
		mcp.WithString("vehicle_type", mcp.Description("To filter listing on their vehicle type")),
		mcp.WithString("ymmt", mcp.Description("Comma separated list of Year, Make, Model, Trim combinations. Each combination needs to have the year,make,model, trim values separated by a pipe '|' character in the form year|make|model|trim. e.g. 2010|Audi|A5,2014|Nissan|Sentra|S 6MT,|Honda|City|   You could just provide strings of the form - 'year|make||' or 'year|make|model' or '|make|model|' combinations. Individual year / make / model filters provied with the API calls will take precedence over the Year, Make, Model, Trim combinations. The Make, Model, Trim values must be valid values as per the Marketcheck Vin Decoder. If you are using a separate vin decoder then look at using the 'vins' or 'taxonomy_vins' parameter to the search api instead the year|make|model|trim combinations.")),
		mcp.WithString("match", mcp.Description("Comma separated list of Year, Make, Model, Trim fields. For example - year,make,model,trim fields for which user wants to do an exact match")),
		mcp.WithString("cylinders", mcp.Description("To filter listing on their cylinders")),
		mcp.WithString("transmission", mcp.Description("To filter listing on their transmission")),
		mcp.WithString("doors", mcp.Description("Doors to filter the cars on. Valid filter values are those that our Search facets API returns for unique doors. You can pass in multiple doors values comma separated")),
		mcp.WithString("drivetrain", mcp.Description("To filter listing on their drivetrain")),
		mcp.WithString("exterior_color", mcp.Description("Exterior color to match. Valid filter values are those that our Search facets API returns for unique exterior colors. You can pass in multiple exterior color values comma separated")),
		mcp.WithString("interior_color", mcp.Description("Interior color to match. Valid filter values are those that our Search facets API returns for unique interior colors. You can pass in multiple interior color values comma separated")),
		mcp.WithString("base_exterior_color", mcp.Description("Base exterior color to match. Valid filter values are those that our Search facets API returns for unique base exterior colors. You can pass in multiple base interior color values comma separated")),
		mcp.WithString("base_interior_color", mcp.Description("Base interior color to match. Valid filter values are those that our Search facets API returns for unique base interior colors. You can pass in multiple base interior color values comma separated")),
		mcp.WithString("engine", mcp.Description("To filter listing on their engine")),
		mcp.WithString("engine_size", mcp.Description("Engine Size to match. Valid filter values are those that our Search facets API returns for unique engine size. You can pass in multiple engine size values comma separated")),
		mcp.WithString("engine_aspiration", mcp.Description("Engine Aspiration to match. Valid filter values are those that our Search facets API returns for unique Engine Aspirations. You can pass in multiple Engine aspirations values comma separated")),
		mcp.WithString("engine_block", mcp.Description("Engine Block to match. Valid filter values are those that our Search facets API returns for unique Engine Block. You can pass in multiple Engine Block values comma separated")),
		mcp.WithString("highway_mpg_range", mcp.Description("Highway mileage range for UK to filter listings with the mileage in the range given. Range to be given in the format - min-max e.g. 1000-5000")),
		mcp.WithString("city_mpg_range", mcp.Description("City mileage range for UK to filter listings with the mileage in the range given. Range to be given in the format - min-max e.g. 1000-5000")),
		mcp.WithString("combined_mpg_range", mcp.Description("Combined mileage range for UK to filter listings with the mileage in the range given. Range to be given in the format - min-max e.g. 1000-5000")),
		mcp.WithString("miles_range", mcp.Description("Miles range to filter listings with miles in the given range. Range to be given in the format - min-max e.g. 1000-5000")),
		mcp.WithString("price_range", mcp.Description("Price range to filter listings with the price in the range given. Range to be given in the format - min-max e.g. 1000-5000")),
		mcp.WithString("msrp_range", mcp.Description("MSRP range to filter listings with the msrp in the range given. Range to be given in the format - min-max e.g. 1000-5000")),
		mcp.WithString("dom_range", mcp.Description("Days on Market range to filter cars with the DOM within the given range. Range to be given in the format - min-max e.g. 10-50")),
		mcp.WithString("last_seen_range", mcp.Description("Last seen date range to filter listings with the last seen in the range given. Range to be given in the format [YYYYMMDD] - min-max e.g. 20190523-20190623")),
		mcp.WithString("first_seen_range", mcp.Description("First seen date range to filter listings with the first seen in the range given. Range to be given in the format [YYYYMMDD] - min-max e.g. 20190523-20190623")),
		mcp.WithString("first_seen_at_source_range", mcp.Description("First seen at source date range to filter listings with the first seen at source in the range given. Range to be given in the format [YYYYMMDD] - min-max e.g. 20190523-20190623")),
		mcp.WithString("first_seen_at_mc_range", mcp.Description("First seen at MC date range to filter listings with the first seen at MC in the range given. Range to be given in the format [YYYYMMDD] - min-max e.g. 20190523-20190623")),
		mcp.WithString("last_seen_days", mcp.Description("Last seen days range to filter listings with the last seen in the range given. Range to be given in the format - min-max e.g. 25-12")),
		mcp.WithString("first_seen_days", mcp.Description("First seen days range to filter listings with the first seen in the range given. Range to be given in the format - min-max e.g. 25-12")),
		mcp.WithString("first_seen_at_source_days", mcp.Description("First seen at source days range to filter listings with the first seen at source in the range given. Range to be given in the format - min-max e.g. 25-12")),
		mcp.WithString("first_seen_at_mc_days", mcp.Description("First seen at MC days range to filter listings with the first seen at MC in the range given. Range to be given in the format - min-max e.g. 25-12")),
		mcp.WithString("sort_by", mcp.Description("Sort by field. Default sort field is distance from the given point")),
		mcp.WithString("sort_order", mcp.Description("Sort order - asc or desc. Default sort order is asc")),
		mcp.WithNumber("rows", mcp.Description("Number of results to return. Default is 10. Max is 50")),
		mcp.WithNumber("start", mcp.Description("Page number to fetch the results for the given criteria. Default is 0. Pagination is allowed only till first 10000 results for the search and sort criteria. The page value can be only between 1 to 10000/rows")),
		mcp.WithBoolean("include_non_vin_listings", mcp.Description("To include non vin listings. Default is false")),
		mcp.WithString("facets", mcp.Description("The comma separated list of fields for which facets are requested. Facets could be requested in addition to the listings for the search. Please note - The API calls with lots of facet fields may take longer to respond.")),
		mcp.WithString("range_facets", mcp.Description("The comma separated list of numeric fields for which range facets are requested. Range facets could be requested in addition to the listings for the search. Please note - The API calls with lots of range facet fields may take longer to respond.")),
		mcp.WithString("facet_sort", mcp.Description("Control sort order of facets with this parameter with default sort being on count, Other available sort is alphabetical sort, which can be obtained by using index as value for this param")),
		mcp.WithString("stats", mcp.Description("The list of fields for which stats need to be generated based on the matching listings for the search criteria. The stats consists of mean, max, average and count of listings based on which the stats are calculated for the field. Stats could be requested in addition to the listings for the search. Please note - The API calls with the stats fields may take longer to respond.")),
		mcp.WithString("country", mcp.Description("To filter listing on Country in which they are listed")),
		mcp.WithBoolean("plot", mcp.Description("If plot has value true results in around 25k coordinates with limited fields to plot respective graph")),
		mcp.WithBoolean("nodedup", mcp.Description("If nodedup is set to true then API will give results without is_searchable i.e multiple listings for single vin")),
		mcp.WithBoolean("dedup", mcp.Description("If dedup is set to true then will give results with is_searchable irrespecive of dealer_id or source")),
		mcp.WithBoolean("owned", mcp.Description("Used in combination with dealer_id or source, when true returns the listings actually owned by dealer himself")),
		mcp.WithString("state", mcp.Description("To filter listing on State in which they are listed")),
		mcp.WithString("city", mcp.Description("To filter listing on City in which they are listed")),
		mcp.WithString("msa_code", mcp.Description("To filter listing on msa code in which they are listed")),
		mcp.WithString("dealer_name", mcp.Description("Filter listings on dealer_name")),
		mcp.WithString("dealership_group_name", mcp.Description("Name of the dealership group to search for")),
		mcp.WithString("trim_o", mcp.Description("Filter listings on web scraped trim")),
		mcp.WithString("trim_r", mcp.Description("Filter trim on custom possible matches")),
		mcp.WithString("dom_active_range", mcp.Description("Active Days on Market range to filter cars with the DOM within the given range. Range to be given in the format - min-max e.g. 10-50")),
		mcp.WithString("dom_180_range", mcp.Description("Last 180 Days on Market range to filter cars with the DOM within the given range. Range to be given in the format - min-max e.g. 10-50")),
		mcp.WithBoolean("exclude_certified", mcp.Description("Boolean param to exclude certified cars from search results")),
		mcp.WithString("fuel_type", mcp.Description("To filter listing on their fuel type")),
		mcp.WithString("dealer_type", mcp.Description("Filter based on dealer type independant or franchise")),
		mcp.WithBoolean("photo_links", mcp.Description("A boolean indicating whether to include only those listings that have photo_links in search results, And discard those that don't have them")),
		mcp.WithBoolean("photo_links_cached", mcp.Description("A boolean indicating whether to include only those listings that have photo_links_cached in search results, And discard those that don't have them")),
		mcp.WithString("stock_no", mcp.Description("To filter listing on their stock number on lot")),
		mcp.WithBoolean("sold", mcp.Description("sold parameter to fetch only sold listings")),
		mcp.WithBoolean("include_relevant_links", mcp.Description("To include_relevant_links. Default is true")),
		mcp.WithString("expired", mcp.Description("Boolean falg to either fetch only the expired listings or active ones")),
		mcp.WithString("exclude_dealer_ids", mcp.Description("A list of dealer ids to exclude from result")),
		mcp.WithString("exclude_sources", mcp.Description("A list of sources to exclude from result")),
		mcp.WithString("in_transit", mcp.Description("A boolean to filter in transit vehicles")),
		mcp.WithString("seating_capacity", mcp.Description("To filter on vehicle seating capacity")),
		mcp.WithString("insurance_group", mcp.Description("Insurance Group")),
		mcp.WithString("vrm", mcp.Description("To filter on vrm")),
		mcp.WithString("num_owners", mcp.Description("Number of owners. Range to be given in the format - min-max e.g. 1000-5000")),
		mcp.WithString("variant", mcp.Description("To filter listing on their variant")),
		mcp.WithString("postal_code", mcp.Description("To filter listing on postal code around which they are listed")),
		mcp.WithString("write_off_category", mcp.Description("write off category")),
		mcp.WithString("fca_status", mcp.Description("To filter on fca status")),
		mcp.WithString("active_inventory_date_range", mcp.Description("date range to filter listings that were active within given date range. Range to be given in the format [YYYYMMDD] - min-max e.g. 20190523-20190623")),
		mcp.WithString("engine_size_range", mcp.Description("Engine size range to filter listings with engine size in the given range. Range to be given in the format - min-max e.g. 1.0-2")),
		mcp.WithString("price_change_range", mcp.Description("Price change range to filter listings with price change within given price_change_range. Range to be given in the format - min-max e.g. 10-500")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_search_car_uk_recentsHandler(cfg),
	}
}

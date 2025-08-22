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

func GetmdsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["api_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("api_key=%v", val))
		}
		if val, ok := args["vin"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("vin=%v", val))
		}
		if val, ok := args["exact"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("exact=%v", val))
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
		if val, ok := args["msa_code"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("msa_code=%v", val))
		}
		if val, ok := args["debug"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("debug=%v", val))
		}
		if val, ok := args["include_sold"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_sold=%v", val))
		}
		if val, ok := args["country"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("country=%v", val))
		}
		if val, ok := args["state"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("state=%v", val))
		}
		if val, ok := args["city"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("city=%v", val))
		}
		if val, ok := args["ymmt"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ymmt=%v", val))
		}
		if val, ok := args["car_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("car_type=%v", val))
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
		if val, ok := args["carfax_1_owner"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("carfax_1_owner=%v", val))
		}
		if val, ok := args["carfax_clean_title"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("carfax_clean_title=%v", val))
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
		if val, ok := args["dealership_group_name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dealership_group_name=%v", val))
		}
		if val, ok := args["dom_active_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dom_active_range=%v", val))
		}
		if val, ok := args["dom_180_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dom_180_range=%v", val))
		}
		if val, ok := args["fuel_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fuel_type=%v", val))
		}
		if val, ok := args["dealer_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dealer_type=%v", val))
		}
		if val, ok := args["engine_size_range"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("engine_size_range=%v", val))
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
		url := fmt.Sprintf("%s/mds/car%s", cfg.BaseURL, queryString)
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
		var result models.Mds
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

func CreateGetmdsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_mds_car",
		mcp.WithDescription("Market Days Supply"),
		mcp.WithString("api_key", mcp.Description("The API Authentication Key. Mandatory with all API calls.")),
		mcp.WithString("vin", mcp.Description("VIN to decode")),
		mcp.WithBoolean("exact", mcp.Description("Exact parameter")),
		mcp.WithString("latitude", mcp.Description("Latitude component of location")),
		mcp.WithString("longitude", mcp.Description("Longitude component of location")),
		mcp.WithNumber("radius", mcp.Description("Radius around the search location (Unit - Miles)")),
		mcp.WithString("zip", mcp.Description("To filter listing on ZIP around which they are listed")),
		mcp.WithString("msa_code", mcp.Description("To filter listing on msa code in which they are listed")),
		mcp.WithBoolean("debug", mcp.Description("Debug parameter")),
		mcp.WithBoolean("include_sold", mcp.Description("To fetch sold vins")),
		mcp.WithString("country", mcp.Description("To filter listing on Country in which they are listed")),
		mcp.WithString("state", mcp.Description("To filter listing on State in which they are listed")),
		mcp.WithString("city", mcp.Description("To filter listing on City in which they are listed")),
		mcp.WithString("ymmt", mcp.Description("Comma separated list of Year, Make, Model, Trim combinations. Each combination needs to have the year,make,model, trim values separated by a pipe '|' character in the form year|make|model|trim. e.g. 2010|Audi|A5,2014|Nissan|Sentra|S 6MT,|Honda|City|   You could just provide strings of the form - 'year|make||' or 'year|make|model' or '|make|model|' combinations. Individual year / make / model filters provied with the API calls will take precedence over the Year, Make, Model, Trim combinations. The Make, Model, Trim values must be valid values as per the Marketcheck Vin Decoder. If you are using a separate vin decoder then look at using the 'vins' or 'taxonomy_vins' parameter to the search api instead the year|make|model|trim combinations.")),
		mcp.WithString("car_type", mcp.Description("Car type. Allowed values are - new / used / certified")),
		mcp.WithString("lease_term", mcp.Description("Search listings with exact lease term, or inside a range with min and max seperated by a dash like lease_term=30-60")),
		mcp.WithString("lease_down_payment", mcp.Description("Search listings with exact down payment in lease offers, or inside a range with min and max seperated by a dash like lease_down_payment=30-60")),
		mcp.WithString("lease_emp", mcp.Description("Search listings with lease offers exactly matching Estimated Monthly Payment(EMI), or inside a range with min and max seperated by a dash like lease_emp=30-60")),
		mcp.WithString("finance_loan_term", mcp.Description("Search listings with exact finance loan term, or inside a range with min and max seperated by a dash like finance_loan_term=30-60")),
		mcp.WithString("finance_loan_apr", mcp.Description("Search listings with finance offers exactly matching loans Annual Percentage Rate, or inside a range with min and max seperated by a dash like finance_loan_apr=30-60")),
		mcp.WithString("finance_emp", mcp.Description("Search listings with finance offers exactly matching Estimated Monthly Payment(EMI), or inside a range with min and max seperated by a dash like finance_emp=30-60")),
		mcp.WithString("finance_down_payment", mcp.Description("Search listings with exact down payment in finance offers, or inside a range with min and max seperated by a dash like finance_down_payment=30-60")),
		mcp.WithString("finance_down_payment_per", mcp.Description("Search listings with exact down payment percentage in finance offers, or inside a range with min and max seperated by a dash like finance_down_payment_per=30-60")),
		mcp.WithString("carfax_1_owner", mcp.Description("Indicates whether car has had only one owner or not")),
		mcp.WithString("carfax_clean_title", mcp.Description("Indicates whether car has clean ownership records")),
		mcp.WithString("year", mcp.Description("To filter listing on their year")),
		mcp.WithString("make", mcp.Description("To filter listings on their make")),
		mcp.WithString("model", mcp.Description("To filter listings on their model")),
		mcp.WithString("trim", mcp.Description("To filter listing on their trim")),
		mcp.WithString("dealer_id", mcp.Description("Dealer id to filter the listings.")),
		mcp.WithString("source", mcp.Description("To filter listing on their source")),
		mcp.WithString("body_type", mcp.Description("To filter listing on their body type")),
		mcp.WithString("body_subtype", mcp.Description("Body subtype to filter the listings on. Valid filter values are those that our Search facets API returns for unique body subtypes. You can pass in multiple body subtype values comma separated")),
		mcp.WithString("vehicle_type", mcp.Description("To filter listing on their vehicle type")),
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
		mcp.WithString("highway_mpg_range", mcp.Description("Highway mileage range to filter listings with the mileage in the range given. Range to be given in the format - min-max e.g. 1000-5000")),
		mcp.WithString("city_mpg_range", mcp.Description("City mileage range to filter listings with the mileage in the range given. Range to be given in the format - min-max e.g. 1000-5000")),
		mcp.WithString("miles_range", mcp.Description("Miles range to filter listings with miles in the given range. Range to be given in the format - min-max e.g. 1000-5000")),
		mcp.WithString("price_range", mcp.Description("Price range to filter listings with the price in the range given. Range to be given in the format - min-max e.g. 1000-5000")),
		mcp.WithString("msrp_range", mcp.Description("MSRP range to filter listings with the msrp in the range given. Range to be given in the format - min-max e.g. 1000-5000")),
		mcp.WithString("dom_range", mcp.Description("Days on Market range to filter cars with the DOM within the given range. Range to be given in the format - min-max e.g. 10-50")),
		mcp.WithString("dealership_group_name", mcp.Description("Name of the dealership group to search for")),
		mcp.WithString("dom_active_range", mcp.Description("Active Days on Market range to filter cars with the DOM within the given range. Range to be given in the format - min-max e.g. 10-50")),
		mcp.WithString("dom_180_range", mcp.Description("Last 180 Days on Market range to filter cars with the DOM within the given range. Range to be given in the format - min-max e.g. 10-50")),
		mcp.WithString("fuel_type", mcp.Description("To filter listing on their fuel type")),
		mcp.WithString("dealer_type", mcp.Description("Filter based on dealer type independant or franchise")),
		mcp.WithString("engine_size_range", mcp.Description("Engine size range to filter listings with engine size in the given range. Range to be given in the format - min-max e.g. 1.0-2")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetmdsHandler(cfg),
	}
}

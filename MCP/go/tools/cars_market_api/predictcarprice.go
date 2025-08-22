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

func PredictcarpriceHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["car_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("car_type=%v", val))
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
		if val, ok := args["is_certified"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_certified=%v", val))
		}
		if val, ok := args["carfax_1_owner"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("carfax_1_owner=%v", val))
		}
		if val, ok := args["carfax_clean_title"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("carfax_clean_title=%v", val))
		}
		if val, ok := args["base_exterior_color"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("base_exterior_color=%v", val))
		}
		if val, ok := args["base_interior_color"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("base_interior_color=%v", val))
		}
		if val, ok := args["transmission"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("transmission=%v", val))
		}
		if val, ok := args["drivetrain"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("drivetrain=%v", val))
		}
		if val, ok := args["engine_size"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("engine_size=%v", val))
		}
		if val, ok := args["engine_block"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("engine_block=%v", val))
		}
		if val, ok := args["cylinders"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cylinders=%v", val))
		}
		if val, ok := args["doors"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("doors=%v", val))
		}
		if val, ok := args["highway_mpg"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("highway_mpg=%v", val))
		}
		if val, ok := args["city_mpg"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("city_mpg=%v", val))
		}
		if val, ok := args["latitude"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("latitude=%v", val))
		}
		if val, ok := args["longitude"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("longitude=%v", val))
		}
		if val, ok := args["miles"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("miles=%v", val))
		}
		if val, ok := args["zip"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("zip=%v", val))
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
		url := fmt.Sprintf("%s/predict/car/price%s", cfg.BaseURL, queryString)
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
		var result models.PricePrediction
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

func CreatePredictcarpriceTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_predict_car_price",
		mcp.WithDescription("Predict car price based on it's specifications"),
		mcp.WithString("api_key", mcp.Description("The API Authentication Key. Mandatory with all API calls.")),
		mcp.WithString("vin", mcp.Description("Predict price for a VIN")),
		mcp.WithString("car_type", mcp.Required(), mcp.Description("Car condition")),
		mcp.WithNumber("year", mcp.Description("Car manufacturing year")),
		mcp.WithString("make", mcp.Description("Car's make")),
		mcp.WithString("model", mcp.Description("Car's model")),
		mcp.WithString("trim", mcp.Description("Car's trim")),
		mcp.WithBoolean("is_certified", mcp.Description("Boolean to indicate car is certified or not")),
		mcp.WithBoolean("carfax_1_owner", mcp.Description("Boolean to indicate car is carfax one owner or not")),
		mcp.WithBoolean("carfax_clean_title", mcp.Description("Boolean to indicate car has clean title or not")),
		mcp.WithString("base_exterior_color", mcp.Description("Base exterior color of the car")),
		mcp.WithString("base_interior_color", mcp.Description("Base interior color of the car")),
		mcp.WithString("transmission", mcp.Description("Transmission on the car")),
		mcp.WithString("drivetrain", mcp.Description("Drivetrain on the car")),
		mcp.WithString("engine_size", mcp.Description("Engine Size of the car")),
		mcp.WithString("engine_block", mcp.Description("Engine Block of the car")),
		mcp.WithNumber("cylinders", mcp.Description("Number of cylinders in the vehicle")),
		mcp.WithNumber("doors", mcp.Description("Number of doors in the vehicle")),
		mcp.WithNumber("highway_mpg", mcp.Description("Highway mileage")),
		mcp.WithNumber("city_mpg", mcp.Description("City mileage of the car")),
		mcp.WithString("latitude", mcp.Description("Latitude component of the location")),
		mcp.WithString("longitude", mcp.Description("Longitude component of the location")),
		mcp.WithNumber("miles", mcp.Description("miles vehicle has driven in total")),
		mcp.WithString("zip", mcp.Description("Location zip")),
		mcp.WithString("country", mcp.Description("Country for which car price will be predicted")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    PredictcarpriceHandler(cfg),
	}
}

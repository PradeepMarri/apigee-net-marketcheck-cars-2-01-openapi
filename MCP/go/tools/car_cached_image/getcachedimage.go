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

func GetcachedimageHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		listingIDVal, ok := args["listingID"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: listingID"), nil
		}
		listingID, ok := listingIDVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: listingID"), nil
		}
		imageIDVal, ok := args["imageID"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: imageID"), nil
		}
		imageID, ok := imageIDVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: imageID"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["api_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("api_key=%v", val))
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
		url := fmt.Sprintf("%s/image/cache/car/%s/%s%s", cfg.BaseURL, listingID, imageID, queryString)
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
		var result string
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

func CreateGetcachedimageTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_image_cache_car_listingID_imageID",
		mcp.WithDescription("Fetch cached image"),
		mcp.WithString("api_key", mcp.Description("The API Authentication Key. Mandatory with all API calls.")),
		mcp.WithString("listingID", mcp.Required(), mcp.Description("ID of the listing to fetch cached images for")),
		mcp.WithString("imageID", mcp.Required(), mcp.Description("ID of the image to fetch")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetcachedimageHandler(cfg),
	}
}

package main

import (
	"github.com/marketcheck-apis/mcp-server/config"
	"github.com/marketcheck-apis/mcp-server/models"
	tools_car_search "github.com/marketcheck-apis/mcp-server/tools/car_search"
	tools_rank_car_listings "github.com/marketcheck-apis/mcp-server/tools/rank_car_listings"
	tools_motorcycle_search "github.com/marketcheck-apis/mcp-server/tools/motorcycle_search"
	tools_dealer_api "github.com/marketcheck-apis/mcp-server/tools/dealer_api"
	tools_rv_search "github.com/marketcheck-apis/mcp-server/tools/rv_search"
	tools_vin_decoder_api "github.com/marketcheck-apis/mcp-server/tools/vin_decoder_api"
	tools_heavy_equipment_search "github.com/marketcheck-apis/mcp-server/tools/heavy_equipment_search"
	tools_cars_history_api "github.com/marketcheck-apis/mcp-server/tools/cars_history_api"
	tools_oem_incentive_search "github.com/marketcheck-apis/mcp-server/tools/oem_incentive_search"
	tools_cars_market_api "github.com/marketcheck-apis/mcp-server/tools/cars_market_api"
	tools_car_cached_image "github.com/marketcheck-apis/mcp-server/tools/car_cached_image"
	tools_crm_cleanse_api "github.com/marketcheck-apis/mcp-server/tools/crm_cleanse_api"
	tools_recall_search "github.com/marketcheck-apis/mcp-server/tools/recall_search"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_car_search.CreateGet_listing_car_fsbo_id_mediaTool(cfg),
		tools_rank_car_listings.CreateSearchandrankcarTool(cfg),
		tools_motorcycle_search.CreateGet_search_motorcycle_auto_completeTool(cfg),
		tools_dealer_api.CreateGet_dealer_car_uk_idTool(cfg),
		tools_rv_search.CreateGet_listing_rv_uk_id_mediaTool(cfg),
		tools_vin_decoder_api.CreateDecodevianeovinTool(cfg),
		tools_motorcycle_search.CreateGet_listing_motorcycle_idTool(cfg),
		tools_motorcycle_search.CreateGet_listing_motorcycle_id_extraTool(cfg),
		tools_dealer_api.CreateGet_dealer_motorcycle_idTool(cfg),
		tools_car_search.CreateGet_listing_car_fsbo_idTool(cfg),
		tools_motorcycle_search.CreateGet_listing_motorcycle_id_mediaTool(cfg),
		tools_dealer_api.CreateGet_dealers_heavy_equipmentTool(cfg),
		tools_heavy_equipment_search.CreateGet_listing_heavy_equipment_idTool(cfg),
		tools_rv_search.CreateGet_search_rv_activeTool(cfg),
		tools_cars_history_api.CreateGet_history_car_uk_vrmTool(cfg),
		tools_oem_incentive_search.CreateOemsearchTool(cfg),
		tools_car_search.CreateGet_listing_car_id_mediaTool(cfg),
		tools_car_search.CreateGet_search_car_auction_activeTool(cfg),
		tools_car_search.CreateGet_listing_car_id_extraTool(cfg),
		tools_car_search.CreateGet_listing_car_uk_id_extraTool(cfg),
		tools_rv_search.CreateGet_listing_rv_id_extraTool(cfg),
		tools_cars_market_api.CreatePredictcarpriceTool(cfg),
		tools_rv_search.CreateGet_search_rv_auto_completeTool(cfg),
		tools_rv_search.CreateGet_listing_rv_idTool(cfg),
		tools_car_search.CreateSearchTool(cfg),
		tools_car_search.CreateGet_listing_car_uk_id_mediaTool(cfg),
		tools_car_search.CreateGetlistingTool(cfg),
		tools_rv_search.CreateGet_listing_rv_uk_idTool(cfg),
		tools_cars_market_api.CreateGetsalescountTool(cfg),
		tools_vin_decoder_api.CreateDecodeTool(cfg),
		tools_heavy_equipment_search.CreateGet_search_heavy_equipment_auto_completeTool(cfg),
		tools_dealer_api.CreateGetdealerTool(cfg),
		tools_vin_decoder_api.CreateGettaxonomytermsTool(cfg),
		tools_car_search.CreateGet_listing_car_auction_id_extraTool(cfg),
		tools_rv_search.CreateGet_listing_rv_uk_id_extraTool(cfg),
		tools_cars_market_api.CreatePredictukcarpriceTool(cfg),
		tools_dealer_api.CreateGet_dealers_rvTool(cfg),
		tools_car_search.CreateGet_listing_car_auction_id_mediaTool(cfg),
		tools_vin_decoder_api.CreateGet_specs_car_auto_completeTool(cfg),
		tools_car_search.CreateGet_search_car_recentsTool(cfg),
		tools_dealer_api.CreateGet_dealers_car_ukTool(cfg),
		tools_car_search.CreateGet_listing_car_auction_idTool(cfg),
		tools_dealer_api.CreateGet_dealer_heavy_equipment_idTool(cfg),
		tools_cars_market_api.CreateFarevalueTool(cfg),
		tools_car_cached_image.CreateGetcachedimageTool(cfg),
		tools_cars_market_api.CreateGetmdsTool(cfg),
		tools_rv_search.CreateGet_search_rv_uk_activeTool(cfg),
		tools_car_search.CreateGet_search_car_uk_recentsTool(cfg),
		tools_car_search.CreateAutocompleteTool(cfg),
		tools_car_search.CreateGet_listing_car_uk_idTool(cfg),
		tools_dealer_api.CreateGet_dealer_rv_idTool(cfg),
		tools_rank_car_listings.CreateRankcarTool(cfg),
		tools_motorcycle_search.CreateGet_search_motorcycle_activeTool(cfg),
		tools_rv_search.CreateGet_listing_rv_id_mediaTool(cfg),
		tools_crm_cleanse_api.CreateCrmcheckTool(cfg),
		tools_car_search.CreateGet_search_car_activeTool(cfg),
		tools_car_search.CreateGet_car_dealer_inventory_activeTool(cfg),
		tools_heavy_equipment_search.CreateGet_listing_heavy_equipment_id_extraTool(cfg),
		tools_heavy_equipment_search.CreateGet_search_heavy_equipment_activeTool(cfg),
		tools_cars_history_api.CreateGetcarhistoryTool(cfg),
		tools_dealer_api.CreateGet_dealers_motorcycleTool(cfg),
		tools_cars_market_api.CreateGetpopularcarsTool(cfg),
		tools_cars_market_api.CreateGetdailystatsTool(cfg),
		tools_heavy_equipment_search.CreateGet_listing_heavy_equipment_id_mediaTool(cfg),
		tools_dealer_api.CreateDealersearchTool(cfg),
		tools_recall_search.CreateGetrecallhistoryTool(cfg),
		tools_car_search.CreateGet_listing_car_fsbo_id_extraTool(cfg),
		tools_car_search.CreateGet_search_car_fsbo_activeTool(cfg),
		tools_vin_decoder_api.CreateDecodeviaepiTool(cfg),
	}
}

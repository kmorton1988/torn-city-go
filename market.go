package torn

import (
	"encoding/json"
	"strconv"
	"strings"
)

// Market represents market listing data
// for a particular item or points.
type Market struct {
	Bazaar map[string]*struct {
		Cost     int `json:"cost"`
		Quantity int `json:"quantity"`
	} `json:"bazaar"`
	ItemMarket map[string]*struct {
		Cost     int `json:"cost"`
		Quantity int `json:"quantity"`
	} `json:"itemmarket"`
	PointsMarket map[string]*struct {
		Cost      int `json:"cost"`
		Quantity  int `json:"quantity"`
		TotalCost int `json:"total_cost"`
	} `json:"pointsmarket"`
	Timestamp int `json:"timestamp"`
}

// GetMarketValue returns data for a specific Torn Item, by ID.
// GetMarketValue supports bazaar data as well as the item market and points market.
// An item ID is not necessary for the points market.
// See https://www.torn.com/api.html
func (s *Session) GetMarketValue(ID int, args ...string) (market *Market, err error) {
	var itemID string
	if ID != 0 {
		itemID = strconv.Itoa(ID)
	}

	var selections string
	for _, arg := range args {
		selections += arg + ","
	}
	selections = strings.TrimSuffix(selections, ",")

	data, err := s.callAPI(apiMarket+endpoint(itemID), map[string]string{"selections": selections})
	if err != nil {
		return
	}
	market = &Market{}
	json.Unmarshal(data, market)
	return
}

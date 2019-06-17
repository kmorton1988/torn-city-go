package torn

import (
	"encoding/json"
	"strconv"
)

type propertyWrapper struct {
	Data *Property `json:"property"`
}

// Property represents all data of a Torn property
type Property struct {
	OwnerID  int      `json:"owner_id"`
	Type     int      `json:"property_type"`
	Happy    int      `json:"happy"`
	Upkeep   int      `json:"upkeep"`
	Upgrades []string `json:"upgrades"`
	Staff    []string `json:"staff"`
	Rented   *struct {
		UserID     string `json:"user_id"`
		DaysLeft   string `json:"days_left"`
		TotalCost  string `json:"total_cost"`
		CostPerDay string `json:"cost_per_day"`
	} `json:"rented"`
	UsersLiving string `json:"users_living"`
}

// QueryProperty returns data for a specific Torn property, by ID.
// See https://www.torn.com/api.html
func (s *Session) QueryProperty(ID int) (property *Property, err error) {
	var propertyID string
	if ID != 0 {
		propertyID = strconv.Itoa(ID)
	}

	data, err := s.callAPI(apiProperties+endpoint(propertyID), nil)
	if err != nil {
		return
	}
	wrapper := &propertyWrapper{}
	json.Unmarshal(data, wrapper)
	property = wrapper.Data
	return
}

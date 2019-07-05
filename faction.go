package torn

import (
	"encoding/json"
	"strconv"
	"strings"
)

// Faction represents all data of a Torn faction
type Faction struct {
	ID        string `json:"ID"`
	Name      string `json:"name"`
	Leader    int    `json:"leader"`
	CoLeader  int    `json:"co-leader"`
	Respect   int    `json:"respect"`
	Age       int    `json:"age"`
	BestChain int    `json:"best_chain"`
	// TODO wars
	Members map[string]*struct {
		Name          string   `json:"name"`
		DaysInFaction int      `json:"days_in_faction"`
		LastAction    string   `json:"last_action"`
		Status        []string `json:"status"`
	} `json:"members"`
	Peace  map[string]int64 `json:"peace"`
	Chains map[string]*struct {
		Chain   int    `json:"chain"`
		Respect string `json:"respect"`
		Start   int64  `json:"start"`
		End     int64  `json:"end"`
	} `json:"chains"`
	Donations map[string]*struct {
		Name          string `json:"name"`
		MoneyBalance  int64  `json:"money_balance"`
		PointsBalance int64  `json:"points_balance"`
	} `json:"donations"`
	// TODO there is a lot left for faction data.
}

// QueryFaction returns data for a specific Torn faction, by ID.
// QueryFaction can take multiple additional options for addition data.
// See https://www.torn.com/api.html
func (s *Session) QueryFaction(ID int, args ...string) (faction *Faction, err error) {
	var factionID string
	if ID != 0 {
		factionID = strconv.Itoa(ID)
	}

	var selections string
	for _, arg := range args {
		selections += arg + ","
	}
	selections = strings.TrimSuffix(selections, ",")

	data, err := s.callAPI(apiUser+endpoint(factionID), map[string]string{"selections": selections})
	if err != nil {
		return
	}
	faction = &Faction{}
	json.Unmarshal(data, faction)
	return
}

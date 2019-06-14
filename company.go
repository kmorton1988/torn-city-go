package torn

import (
	"encoding/json"
	"strconv"
	"strings"
)

type companyWrapper struct {
	Data *Company `json:"company"`
}

// Company represents all data of a Torn company
type Company struct {
	ID              int    `json:"ID"`
	Type            int    `json:"company_type"`
	Rating          int    `json:"rating"`
	Name            string `json:"name"`
	DirectorID      int    `json:"director"`
	DailyProfit     int    `json:"daily_profit"`
	WeeklyProfit    int    `json:"weekly_profit"`
	DailyCustomers  int    `json:"daily_customers"`
	WeeklyCustomers int    `json:"weekly_customers"`
	DaysOld         int    `json:"days_old"`
	Employees       map[string]*struct {
		Name          string `json:"name"`
		Position      string `json:"position"`
		DaysInCompany int    `json:"days_in_company"`
	} `json:"employees"`
	Detailed *struct {
		ID                int `json:"ID"`
		Bank              int `json:"company_bank"`
		Popularity        int `json:"popularity"`
		Efficiency        int `json:"efficiency"`
		Environment       int `json:"environment"`
		TrainsAvailable   int `json:"trains_available"`
		AdvertisingBudget int `json:"advertising_budget"`
		Upgrades          struct {
			CompanySize   int    `json:"company_size"`
			StaffroomSize string `json:"staffroom_size"`
			StorageSize   string `json:"storage_size"`
			StorageSpace  int    `json:"storage_space"`
		} `json:"upgrades"`
	}
	Stock map[string]*struct {
		Cost       int `json:"cost"`
		RRP        int `json:"rrp"`
		Price      int `json:"price"`
		InStock    int `json:"in_stock"`
		OnOrder    int `json:"on_order"`
		SoldAmount int `json:"sold_amount"`
		SoldWorth  int `json:"sold_worth"`
	} `json:"company_stock"`
	EmployeesDetailed map[string]*struct {
		Position      string `json:"position"`
		DaysEmployed  int    `json:"days_employed"`
		Wage          int    `json:"wage"`
		Effectiveness int    `json:"effectiveness"`
		ManualLabor   int    `json:"manual_labor"`
		Intelligence  int    `json:"intelligence"`
		Endurance     int    `json:"endurance"`
	} `json:"company_employees"`
	News map[string]*struct {
		Timestamp int    `json:"timestamp"`
		News      string `json:"news"`
	} `json:"news"`
}

// GetCompany returns data for a specific Torn company, by ID.
// GetCompany can take multiple additional options for addition data.
// See https://www.torn.com/api.html
func (s *Session) GetCompany(ID int, args ...string) (company *Company, err error) {
	var companyID string
	if ID != 0 {
		companyID = strconv.Itoa(ID)
	}

	var selections string
	for _, arg := range args {
		selections += arg + ","
	}
	selections = strings.TrimSuffix(selections, ",")

	data, err := s.callAPI(apiCompany+endpoint(companyID), map[string]string{"selections": selections})
	if err != nil {
		return
	}
	wrapper := &companyWrapper{}
	json.Unmarshal(data, wrapper)
	company = wrapper.Data
	return
}

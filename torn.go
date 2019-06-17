package torn

import (
	"encoding/json"
	"strconv"
	"strings"
)

// Torn contains all data found under
// the general "Torn" endpoint of the Torn API.
// This includes items, medals, honors, and other information
// about the game.
type Torn struct {
	Items map[string]*struct {
		Name        string `json:"name"`
		Desc        string `json:"description"`
		Effect      string `json:"effect"`
		Requirement string `json:"requirement"`
		Type        string `json:"type"`
		BuyPrice    int    `json:"buy_price"`
		SellPrice   int    `json:"sell_price"`
		MarketValue int    `json:"market_value"`
		Circulation int    `json:"circulation"`
		Image       string `json:"image"`
	} `json:"items"`
	Medals map[string]*struct {
		Name string `json:"name"`
		Desc string `json:"description"`
		Type string `json:"type"`
	} `json:"medals"`
	Honors map[string]*struct {
		Name        string `json:"name"`
		Desc        string `json:"description"`
		Type        int    `json:"type"`
		Circulation int    `json:"circulation"`
		Rarity      string `json:"rarity"`
	} `json:"honors"`
	OrganisedCrimes map[string]*struct {
		Name       string `json:"name"`
		Members    int    `json:"members"`
		Time       int    `json:"time"`
		MinCash    int    `json:"min_cash"`
		MaxCash    int    `json:"max_cash"`
		MinRespect int    `json:"min_respect"`
		MaxRespect int    `json:"max_respect"`
	} `json:"organisedcrimes"`
	Gyms map[string]*struct {
		Name      string `json:"name"`
		Stage     int    `json:"stage"`
		Cost      int    `json:"cost"`
		Energy    int    `json:"energy"`
		Strength  int    `json:"strength"`
		Speed     int    `json:"speed"`
		Defense   int    `json:"defense"`
		Dexterity int    `json:"dexterity"`
		Note      string `json:"note"`
	} `json:"gyms"`
	Companies map[string]*struct {
		Name             string `json:"name"`
		Cost             int    `json:"cost"`
		DefaultEmployees int    `json:"default_employees"`
		Positions        map[string]*struct {
			ManRequired    int    `json:"man_required"`
			IntRequired    int    `json:"int_required"`
			EndRequired    int    `json:"end_required"`
			ManGain        int    `json:"man_gain"`
			IntGain        int    `json:"int_gain"`
			EndGain        int    `json:"end_gain"`
			SpecialAbility string `json:"special_ability"`
			Desc           string `json:"descriptions"`
		} `json:"positions"`
		Stock map[string]*struct {
			Cost string `json:"cost"`
			RRP  int    `json:"rrp"`
		} `json:"stock"`
		Specials map[string]*struct {
			Effect         string `json:"effect"`
			Cost           int    `json:"cost"`
			RatingRequired int    `json:"rating_required"`
		} `json:"specials"`
	} `json:"companies"`
	Properties map[string]*struct {
		Name              string   `json:"name"`
		Cost              int      `json:"cost"`
		Happy             int      `json:"happy"`
		Upkeep            int      `json:"upkeep"`
		UpgradesAvailable []string `json:"upgrades_available"`
		StaffAvailable    []string `json:"staff_available"`
	} `json:"properties"`
	Education map[string]*struct {
		Name      string `json:"name"`
		Desc      string `json:"description"`
		MoneyCost int    `json:"money_cost"`
		Tier      int    `json:"tier"`
		Duration  int    `json:"duration"`
		Results   struct {
			Perk         []string `json:"perk"`
			ManualLabor  []string `json:"manual_labor"`
			Intelligence []string `json:"intelligence"`
			Endurance    []string `json:"endurance"`
		}
		Prerequisites []int `json:"prerequisites"`
	} `json:"education"`
	Stats *struct {
		Timestamp                     int    `json:"timestamp"`
		UsersTotal                    int    `json:"users_total"`
		UsersMale                     int    `json:"users_male"`
		UsersFemale                   int    `json:"users_female"`
		UsersMarriedcouples           int    `json:"users_marriedcouples"`
		UsersDaily                    int    `json:"users_daily"`
		TotalUsersLogins              int    `json:"total_users_logins"`
		TotalUsersPlaytime            string `json:"total_users_playtime"`
		JobArmy                       int    `json:"job_army"`
		JobGrocer                     int    `json:"job_grocer"`
		JobMedical                    int    `json:"job_medical"`
		JobCasino                     int    `json:"job_casino"`
		JobEducation                  int    `json:"job_education"`
		JobLaw                        int    `json:"job_law"`
		JobCompany                    int    `json:"job_company"`
		JobNone                       int    `json:"job_none"`
		Crimes                        int    `json:"crimes"`
		Jailed                        int    `json:"jailed"`
		MoneyOnhand                   int    `json:"money_onhand"`
		MoneyAverage                  int    `json:"money_average"`
		MoneyCitybank                 int    `json:"money_citybank"`
		Items                         int    `json:"items"`
		Events                        int    `json:"events"`
		PointsTotal                   int    `json:"points_total"`
		PointsMarket                  int    `json:"points_market"`
		PointsAveragecost             int    `json:"points_averagecost"`
		PointsBought                  int    `json:"points_bought"`
		TotalPointsBoughttotal        int    `json:"total_points_boughttotal"`
		TotalAttacksWon               int    `json:"total_attacks_won"`
		TotalAttacksLost              int    `json:"total_attacks_lost"`
		TotalAttacksStalemated        int    `json:"total_attacks_stalemated"`
		TotalAttacksRunaway           int    `json:"total_attacks_runaway"`
		TotalAttacksHits              int    `json:"total_attacks_hits"`
		TotalAttacksMisses            int    `json:"total_attacks_misses"`
		TotalAttacksCriticalhits      int    `json:"total_attacks_criticalhits"`
		TotalAttacksRoundsfired       int    `json:"total_attacks_roundsfired"`
		TotalAttacksStealthed         int    `json:"total_attacks_stealthed"`
		TotalAttacksMoneymugged       int    `json:"total_attacks_moneymugged"`
		TotalAttacksRespectgained     int    `json:"total_attacks_respectgained"`
		TotalItemsMarketbought        int    `json:"total_items_marketbought"`
		TotalItemsBazaarbought        int    `json:"total_items_bazaarbought"`
		TotalItemsAuctionswon         int    `json:"total_items_auctionswon"`
		TotalItemsSent                int    `json:"total_items_sent"`
		TotalTrades                   int    `json:"total_trades"`
		TotalItemsBazaarincome        int    `json:"total_items_bazaarincome"`
		TotalItemsCityfinds           int    `json:"total_items_cityfinds"`
		TotalItemsDumpfinds           int    `json:"total_items_dumpfinds"`
		TotalItemsDumped              int    `json:"total_items_dumped"`
		TotalJailJailed               int    `json:"total_jail_jailed"`
		TotalJailBusted               int    `json:"total_jail_busted"`
		TotalJailBusts                int    `json:"total_jail_busts"`
		TotalJailBailed               int    `json:"total_jail_bailed"`
		TotalJailBailcosts            int    `json:"total_jail_bailcosts"`
		TotalHospitalTrips            int    `json:"total_hospital_trips"`
		TotalHospitalMedicalitemsused int    `json:"total_hospital_medicalitemsused"`
		TotalHospitalRevived          int    `json:"total_hospital_revived"`
		TotalMailsSent                int    `json:"total_mails_sent"`
		TotalMailsSentFriends         int    `json:"total_mails_sent_friends"`
		TotalMailsSentFaction         int    `json:"total_mails_sent_faction"`
		TotalMailsSentCompany         int    `json:"total_mails_sent_company"`
		TotalMailsSentSpouse          int    `json:"total_mails_sent_spouse"`
		TotalClassifiedadsPlaced      int    `json:"total_classifiedads_placed"`
		TotalBountyPlaced             int    `json:"total_bounty_placed"`
		TotalBountyRewards            int    `json:"total_bounty_rewards"`
		TotalTravelAll                int    `json:"total_travel_all"`
		TotalTravelArgentina          int    `json:"total_travel_argentina"`
		TotalTravelMexico             int    `json:"total_travel_mexico"`
		TotalTravelDubai              int    `json:"total_travel_dubai"`
		TotalTravelHawaii             int    `json:"total_travel_hawaii"`
		TotalTravelJapan              int    `json:"total_travel_japan"`
		TotalTravelUnitedkingdom      int    `json:"total_travel_unitedkingdom"`
		TotalTravelSouthafrica        int    `json:"total_travel_southafrica"`
		TotalTravelSwitzerland        int    `json:"total_travel_switzerland"`
		TotalTravelChina              int    `json:"total_travel_china"`
		TotalTravelCanada             int    `json:"total_travel_canada"`
		TotalTravelCaymanislands      int    `json:"total_travel_caymanislands"`
		TotalDrugsUsed                int    `json:"total_drugs_used"`
		TotalDrugsOverdosed           int    `json:"total_drugs_overdosed"`
		TotalDrugsCannabis            int    `json:"total_drugs_cannabis"`
		TotalDrugsEcstacy             int    `json:"total_drugs_ecstacy"`
		TotalDrugsKetamine            int    `json:"total_drugs_ketamine"`
		TotalDrugsLsd                 int    `json:"total_drugs_lsd"`
		TotalDrugsOpium               int    `json:"total_drugs_opium"`
		TotalDrugsShrooms             int    `json:"total_drugs_shrooms"`
		TotalDrugsSpeed               int    `json:"total_drugs_speed"`
		TotalDrugsPcp                 int    `json:"total_drugs_pcp"`
		TotalDrugsXanax               int    `json:"total_drugs_xanax"`
		TotalDrugsVicodin             int    `json:"total_drugs_vicodin"`
		TotalMeritsBought             int    `json:"total_merits_bought"`
		TotalRefillsBought            int    `json:"total_refills_bought"`
		TotalCompanyTrains            int    `json:"total_company_trains"`
		TotalStatenhancersUsed        int    `json:"total_statenhancers_used"`
	} `json:"stats"`
	Stocks map[string]*struct {
		Name            string `json:"name"`
		Acronym         string `json:"acronym"`
		Director        string `json:"director"`
		CurrentPrice    string `json:"current_price"`
		MarketCap       int    `json:"market_cap"`
		TotalShares     int    `json:"total_shares"`
		AvailableShares int    `json:"available_shares"`
		Forecast        string `json:"forecast"`
		Demand          string `json:"demand"`
	} `json:"stocks"`
	FactionTree map[string]map[string]*struct {
		Branch    string `json:"branch"`
		Name      string `json:"name"`
		Ability   string `json:"ability"`
		Challenge string `json:"challenge"`
		BaseCost  int    `json:"base_cost"`
	} `json:"factiontree"`
	Timestamp int `json:"timestamp"`
}

// QueryTorn returns global Torn data, such as items and faction trees.
// QueryTorn can take multiple IDs for various elements, as well as extra
// arguments for additional data.
// See https://www.torn.com/api.html
func (s *Session) QueryTorn(ID int, args ...string) (tornData *Torn, err error) {
	var userID string
	if ID != 0 {
		userID = strconv.Itoa(ID)
	}

	var selections string
	for _, arg := range args {
		selections += arg + ","
	}
	selections = strings.TrimSuffix(selections, ",")

	data, err := s.callAPI(apiTorn+endpoint(userID), map[string]string{"selections": selections})
	if err != nil {
		return
	}
	tornData = &Torn{}
	json.Unmarshal(data, tornData)
	return
}

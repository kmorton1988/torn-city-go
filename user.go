package torn

import (
	"encoding/json"
	"strconv"
	"strings"
)

// User represents all data of a Torn user
type User struct {
	Rank       string   `json:"rank"`
	Level      int      `json:"level"`
	Gender     string   `json:"gender"`
	Property   string   `json:"property"`
	Signup     string   `json:"signup"`
	Awards     int      `json:"awards"`
	Friends    int      `json:"friends"`
	Enemies    int      `json:"enemies"`
	ForumPosts int      `json:"forum_posts"`
	Karma      int      `json:"karma"`
	Age        int      `json:"age"`
	Role       string   `json:"role"`
	Donator    int      `json:"donator"`
	ID         int      `json:"player_id"`
	Name       string   `json:"name"`
	PropertyID int      `json:"property_id"`
	Status     []string `json:"status"`
	Job        *struct {
		Position    string `json:"position"`
		CompanyID   int    `json:"company_id"`
		CompanyName string `json:"company_name"`
	} `json:"job"`
	Faction *struct {
		Position      string `json:"position"`
		ID            int    `json:"faction_id"`
		DaysInFaction int    `json:"days_in_faction"`
		Name          string `json:"faction_name"`
	} `json:"faction"`
	Married *struct {
		SpouseID   int    `json:"spouse_id"`
		SpouseName string `json:"spouse_name"`
		Duration   int    `json:"duration"`
	} `json:"married"`
	States *struct {
		HospitalTimestamp int `json:"hospital_timestamp"`
		JailTimestamp     int `json:"jail_timestamp"`
	}
	LastAction *struct {
		Timestamp int    `json:"timestamp"`
		Relative  string `json:"relative"`
	}
	Networth *struct {
		Pending      int     `json:"pending"`
		Wallet       int     `json:"wallet"`
		Bank         int     `json:"bank"`
		Points       int     `json:"points"`
		Cayman       int     `json:"cayman"`
		Vault        int     `json:"vault"`
		PiggyBank    int     `json:"piggybank"`
		Items        int     `json:"items"`
		DisplayCase  int     `json:"displaycase"`
		Bazaar       int     `json:"bazaar"`
		Properties   int     `json:"properties"`
		StockMarket  int     `json:"stockmarket"`
		AuctionHouse int     `json:"auctionhouse"`
		Company      int     `json:"company"`
		Bookie       int     `json:"bookie"`
		Loan         int     `json:"loan"`
		UnpaidFees   int     `json:"unpaidfees"`
		Total        int     `json:"total"`
		ParseTime    float64 `json:"parsetime"`
	} `json:"networth"`
	Bazaar []*struct {
		ID          int    `json:"ID"`
		Name        string `json:"name"`
		Type        string `json:"type"`
		Quantity    int    `json:"quantity"`
		Price       int    `json:"price"`
		MarketPrice int    `json:"market_price"`
	} `json:"bazaar"`
	Display []*struct {
	} `json:"display"` //TODO
	Inventory []*struct {
		ID          int    `json:"ID"`
		Name        string `json:"name"`
		Type        string `json:"type"`
		Quantity    string `json:"quantity"`
		Equipped    int    `json:"equipped"`
		MarketPrice int    `json:"market_price"`
	} `json:"inventory"`
	HOF *struct {
		Attacks struct {
			Value int `json:"value"`
			Rank  int `json:"rank"`
		} `json:"attacks"`
		BattleStats struct {
			Value int `json:"value"`
			Rank  int `json:"rank"`
		} `json:"battlestats"`
		Busts struct {
			Value int `json:"value"`
			Rank  int `json:"rank"`
		} `json:"busts"`
		Defends struct {
			Value int `json:"value"`
			Rank  int `json:"rank"`
		} `json:"defends"`
		Networth struct {
			Value int `json:"value"`
			Rank  int `json:"rank"`
		} `json:"networth"`
		Offences struct {
			Value int `json:"value"`
			Rank  int `json:"rank"`
		} `json:"offences"`
		Revives struct {
			Value int `json:"value"`
			Rank  int `json:"rank"`
		} `json:"revives"`
		Traveled struct {
			Value int `json:"value"`
			Rank  int `json:"rank"`
		} `json:"traveled"`
		WorkStats struct {
			Value int `json:"value"`
			Rank  int `json:"rank"`
		} `json:"workstats"`
		Level struct {
			Value int `json:"value"`
			Rank  int `json:"rank"`
		} `json:"level"`
		Rank struct {
			Value int `json:"value"`
			Rank  int `json:"rank"`
		} `json:"rank"`
		Respect struct {
			Value int `json:"value"`
			Rank  int `json:"rank"`
		} `json:"respect"`
	} `json:"halloffame"`
	Travel *struct {
		Destination string `json:"destination"`
		Timestamp   int    `json:"timestamp"`
		Departed    int    `json:"departed"`
		TimeLeft    int    `json:"time_left"`
	} `json:"travel"`
	Events map[int]*struct {
		Timestamp int    `json:"timestamp"`
		Event     string `json:"event"`
		Seen      int    `json:"seen"`
	} `json:"events"`
	Messages map[int]*struct {
		Timestamp int    `json:"timestamp"`
		ID        int    `json:"ID"`
		Name      string `json:"name"`
		Type      string `json:"type"`
		Title     string `json:"title"`
		Seen      bool   `json:"seen"`
		Read      bool   `json:"read"`
	} `json:"messages"`
	EduCurrent    int   `json:"education_current"`
	EduTimeleft   int   `json:"education_timeleft"`
	EduCompleted  []int `json:"education_completed"`
	MedalsAwarded []int `json:"medals_awarded"`
	MedalsTime    []int `json:"medals_time"`
	HonorsAwards  []int `json:"honors_awarded"`
	HonorsTime    []int `json:"honors_time"`
	Notifications *struct {
		Messages    int `json:"messages"`
		Events      int `json:"events"`
		Awards      int `json:"awards"`
		Competition int `json:"competition"`
	} `json:"notifications"`
	PersonalStats *struct {
		Networth               int `json:"networth"`
		Logins                 int `json:"logins"`
		UserActivity           int `json:"useractivity"`
		Hospital               int `json:"hospital"`
		DefendsLost            int `json:"defendslost"`
		Killstreak             int `json:"killstreak"`
		MissionCreditsEarned   int `json:"missioncreditsearned"`
		AttackDamage           int `json:"attackdamage"`
		BestDamage             int `json:"bestdamage"`
		AttackHits             int `json:"attackhits"`
		PieHits                int `json:"piehits"`
		AttacksStealthed       int `json:"attacksstealthed"`
		MoneyMugged            int `json:"moneymugged"`
		LargestMugged          int `json:"largestmug"`
		AttacksWon             int `json:"attackswon"`
		HighestBeaten          int `json:"highestbeaten"`
		BestKillStreak         int `json:"bestkillstreak"`
		RoundsFired            int `json:"roundsfired"`
		AttackCriticalHits     int `json:"attackcriticalhits"`
		OneHitKills            int `json:"onehitkills"`
		RifHits                int `json:"rifhits"`
		AttackMisses           int `json:"attackmisses"`
		AttacksLost            int `json:"attackslost"`
		ItemsBought            int `json:"itemsbought"`
		DefendsWon             int `json:"defendswon"`
		MoneyInvested          int `json:"moneyinvested"`
		InvestedProfit         int `json:"investedprofit"`
		RevivesReceived        int `json:"revivesreceived"`
		DumpSearches           int `json:"dumpsearches"`
		DumpFinds              int `json:"dumpfinds"`
		Jailed                 int `json:"jailed"`
		RespectForFaction      int `json:"respectforfaction"`
		ConsumablesUsed        int `json:"consumablesused"`
		AlcoholUsed            int `json:"alcoholused"`
		ItemsDumped            int `json:"itemsdumped"`
		DaysBeenDonator        int `json:"daysbeendonator"`
		WeaponsBought          int `json:"weaponsbought"`
		ContractsCompleted     int `json:"contractscompleted"`
		DukeContractsCompleted int `json:"dukecontractscompleted"`
		OrganisedCrimes        int `json:"organisedcrimes"`
		MedicalItemsUsed       int `json:"medicalitemsused"`
		DrugsUsed              int `json:"drugsused"`
		XanTaken               int `json:"xantaken"`
		ItemsSent              int `json:"itemssent"`
		PisHits                int `json:"pishits"`
		YouRunAway             int `json:"yourunaway"`
		GreHits                int `json:"grehits"`
		BountiesPlaced         int `json:"bountiesplaced"`
		TotalBountySpent       int `json:"totalbountyspent"`
		PointsSold             int `json:"pointssold"`
		LsdTaken               int `json:"lsdtaken"`
		MissionsCompleted      int `json:"missionscompleted"`
		Trades                 int `json:"trades"`
		H2hHits                int `json:"h2hhits"`
		CandyUsed              int `json:"candyused"`
		TerritoryTime          int `json:"territorytime"`
		CityFinds              int `json:"cityfinds"`
		BountiesReceived       int `json:"bountiesreceived"`
		TrainsReceived         int `json:"trainsreceived"`
		TravelTimes            int `json:"traveltimes"`
		TravelTime             int `json:"traveltime"`
		LonTravel              int `json:"lontravel"`
		ItemsBoughtAbroad      int `json:"itemsboughtabroad"`
		MexTravel              int `json:"mextravel"`
		CayTravel              int `json:"caytravel"`
		ArgTravel              int `json:"argtravel"`
		CanTravel              int `json:"cantravel"`
		SouTravel              int `json:"soutravel"`
		SwiTravel              int `json:"switravel"`
		AxeHits                int `json:"axehits"`
		JapTravel              int `json:"japtravel"`
		MailsSent              int `json:"mailssent"`
		Overdosed              int `json:"overdosed"`
		ChiTravel              int `json:"chitravel"`
		DubTravel              int `json:"dubtravel"`
		HawTravel              int `json:"hawtravel"`
		BloodWithdrawn         int `json:"bloodwithdrawn"`
		BoostersUsed           int `json:"boostersused"`
		AttacksAssisted        int `json:"attacksassisted"`
		Rehabs                 int `json:"rehabs"`
		RehabCost              int `json:"rehabcost"`
		VirusesCoded           int `json:"virusescoded"`
		BazaarCustomers        int `json:"bazaarcustomers"`
		BazaarSales            int `json:"bazaarsales"`
		BazaarProfit           int `json:"bazaarprofit"`
		EnergyDrinkUsed        int `json:"energydrinkused"`
		ExtTaken               int `json:"exttaken"`
		PointsBought           int `json:"pointsbought"`
		BooksRead              int `json:"booksread"`
		BountiesCollected      int `json:"bountiescollected"`
		TotalBountyReward      int `json:"totalbountyreward"`
		FriendMailsSent        int `json:"friendmailssent"`
		FactionMailsSent       int `json:"factionmailssent"`
		AttacksDraw            int `json:"attacksdraw"`
		Awards                 int `json:"awards"`
		AuctionsWon            int `json:"auctionswon"`
		AuctionSells           int `json:"auctionsells"`
		MeritsBought           int `json:"meritsbought"`
		Refills                int `json:"refills"`
	} `json:"personalstats"`
	ManualLabor    int `json:"manual_labor"`
	Intelligence   int `json:"intelligence"`
	Endurance      int `json:"endurance"`
	CriminalRecord *struct {
		SellingIllegalProducts int `json:"selling_illegal_products"`
		Theft                  int `json:"theft"`
		AutoTheft              int `json:"auto_theft"`
		DrugDeals              int `json:"drug_deals"`
		ComputerCrimes         int `json:"computer_crimes"`
		Murder                 int `json:"murder"`
		FraudCrimes            int `json:"fraud_crimes"`
		Other                  int `json:"other"`
		Total                  int `json:"total"`
	} `json:"criminalrecord"`
	Icons     map[string]string `json:"icons"`
	Cooldowns *struct {
		Drug    int `json:"drug"`
		Medical int `json:"medical"`
		Booster int `json:"booster"`
	} `json:"cooldowns"`
	Points      int `json:"points"`
	CaymanBank  int `json:"cayman_bank"`
	VaultAmount int `json:"vault_amount"`
	// Networth int `json:"networth"` - duplicate key entry
	JobPerks         []string `json:"job_perks"`
	PropertyPerks    []string `json:"property_perks"`
	StockPerks       []string `json:"stock_perks"`
	MeritPerks       []string `json:"merit_perks"`
	EduPerks         []string `json:"education_perks"`
	EnhancerPerks    []string `json:"enhancer_perks"`
	CompanyPerks     []string `json:"company_perks"`
	FactionPerks     []string `json:"faction_perks"`
	Strength         string   `json:"strength"`
	Speed            string   `json:"speed"`
	Dexterity        string   `json:"dexterity"`
	Defense          string   `json:"defense"`
	BattleStatsTotal string   `json:"total"`
	StrengthMod      float32  `json:"strength_modifier"`
	DefenseMod       float32  `json:"defense_modifier"`
	SpeedMod         float32  `json:"speed_modifier"`
	DexterityMod     float32  `json:"dexterity_modifier"`
	StrengthInfo     []string `json:"strength_info"`
	DefenseInfo      []string `json:"defense_info"`
	SpeedInfo        []string `json:"speed_info"`
	DexterityInfo    []string `json:"dexterity_info"`
	ServerTime       int      `json:"server_time"`
	Happy            *struct {
		Current   int `json:"current"`
		Max       int `json:"maximum"`
		Increment int `json:"increment"`
		Interval  int `json:"interval"`
		TickTime  int `json:"ticktime"`
		FullTime  int `json:"fulltime"`
	} `json:"happy"`
	Life *struct {
		Current   int `json:"current"`
		Max       int `json:"maximum"`
		Increment int `json:"increment"`
		Interval  int `json:"interval"`
		TickTime  int `json:"ticktime"`
		FullTime  int `json:"fulltime"`
	} `json:"life"`
	Energy *struct {
		Current   int `json:"current"`
		Max       int `json:"maximum"`
		Increment int `json:"increment"`
		Interval  int `json:"interval"`
		TickTime  int `json:"ticktime"`
		FullTime  int `json:"fulltime"`
	} `json:"energy"`
	Nerve *struct {
		Current   int `json:"current"`
		Max       int `json:"maximum"`
		Increment int `json:"increment"`
		Interval  int `json:"interval"`
		TickTime  int `json:"ticktime"`
		FullTime  int `json:"fulltime"`
	} `json:"nerve"`
	Chain *struct {
		Current   int `json:"current"`
		Max       int `json:"maximum"`
		Increment int `json:"increment"`
		Interval  int `json:"interval"`
		TickTime  int `json:"ticktime"`
		FullTime  int `json:"fulltime"`
	} `json:"chain"`
	AttackDetails map[string]*struct {
		TimestampStarted    int    `json:"timestamp_started"`
		TimestampEnded      int    `json:"timestamp_ended"`
		AttackerID          int    `json:"attacker_id"`
		AttackerName        string `json:"attacker_name"`
		AttackerFactionID   int    `json:"attacker_faction"`
		AttackerFactionName string `json:"attacker_factionname"`
		DefenderID          int    `json:"defender_id"`
		DefenderName        string `json:"defender_name"`
		DefenderFactionID   int    `json:"defender_faction"`
		DefenderFactionName int    `json:"defender_factionname"`
		Result              string `json:"result"`
		Stealthed           int    `json:"stealthed"`
		RespectGain         int    `json:"respect_gain"`
		Chain               int    `json:"chain"`
		Modifiers           struct {
			FairFight   int    `json:"fairFight"`
			War         int    `json:"war"`
			Retaliation int    `json:"retaliation"`
			GroupAttack int    `json:"groupAttack"`
			Overseas    int    `json:"overseas"`
			ChainBonus  string `json:"chainBonus"`
		}
	} `json:"attacks"`
	Revives map[string]*struct {
	} `json:"revives"` // TODO
	Stocks map[string]*struct {
		ID          int    `json:"stock_id"`
		Shares      int    `json:"shares"`
		BoughtPrice string `json:"bought_price"`
		TimeBought  int    `json:"time_bought"`
		TimeListed  int    `json:"time_listed"`
	} `json:"stocks"`
	Properties map[string]*struct {
		OwnerID       int    `json:"owner_id"`
		Type          int    `json:"property_type"`
		Name          string `json:"property"`
		Status        string `json:"status"`
		Happy         int    `json:"happy"`
		Upkeep        int    `json:"upkeep"`
		StaffCost     int    `json:"staff_cost"`
		MarketPrice   int    `json:"market_price"`
		Modifications struct {
			Interior        int `json:"interior"`
			HotTub          int `json:"hot_tub"`
			Sauna           int `json:"sauna"`
			Pool            int `json:"pool"`
			OpenBar         int `json:"open_bar"`
			ShootingRange   int `json:"shooting_range"`
			Vault           int `json:"vault"`
			MedicalFacility int `json:"medical_facility"`
			Airstrip        int `json:"airstrip"`
			Yacht           int `json:"yacht"`
		} `json:"modifications"`
		Staff struct {
			Maid   int `json:"maid"`
			Guard  int `json:"guard"`
			Pilot  int `json:"pilot"`
			Butler int `json:"butler"`
			Doctor int `json:"doctor"`
		} `json:"staff"`
	} `json:"properties"`
	JobPoints *struct {
		Jobs struct {
			Army      int `json:"army"`
			Medical   int `json:"medical"`
			Casino    int `json:"casino"`
			Education int `json:"education"`
			Law       int `json:"law"`
			Grocer    int `json:"grocer"`
		} `json:"jobs"`
	} `json:"jobpoints"`
	Merits *struct {
		NerveBar              int `json:"Nerve Bar"`
		CriticalHitRate       int `json:"Critical Hit Rate"`
		LifePoints            int `json:"Life Points"`
		CrimeExperience       int `json:"Crime Experience"`
		EduLength             int `json:"Education Length"`
		Awareness             int `json:"Awareness"`
		BankInterest          int `json:"Bank Interest"`
		MasterfulLooting      int `json:"Masterful Looting"`
		Stealth               int `json:"Stealth"`
		Hospitalizing         int `json:"Hospitalizing"`
		Brawn                 int `json:"Brawn"`
		Protection            int `json:"Protection"`
		Sharpness             int `json:"Sharpness"`
		Evasion               int `json:"Evasion"`
		HeavyArtilleryMastery int `json:"Heavy Artillery Mastery"`
		MachineGunMastery     int `json:"Machine Gun Mastery"`
		RifleMastery          int `json:"Rifle Mastery"`
		SMGMastery            int `json:"SMG Mastery"`
		ShotgunMastery        int `json:"Shotgun Mastery"`
		PistolMastery         int `json:"Pistol Master"`
		ClubMastery           int `json:"Club Mastery"`
		PiercingMastery       int `json:"Piercing Mastery"`
		SlashingMastery       int `json:"Slashing Mastery"`
		MechanicalMastery     int `json:"Mechanical Mastery"`
		TemporaryMastery      int `json:"Temporary Mastery"`
	} `json:"merits"`
	Refills *struct {
		EnergyRefillUsed        bool `json:"energy_refill_used"`
		NerveRefillUsed         bool `json:"nerve_refill_used"`
		TokenRefillUsed         bool `json:"token_refill_used"`
		SpecialRefillsAvailable int  `json:"special_refills_available"`
	} `json:"refills"`
	Discord *struct {
		UserID    int    `json:"userID"`
		DiscordID string `json:"discordID"`
	} `json:"discord"`
	ActiveGym int `json:"active_gym"`
	Timestamp int `json:"timestamp"`
}

// GetUser returns data for a specific Torn user, by ID.
// GetUser can take multiple additional options for addition data.
// See https://www.torn.com/api.html
func (s *Session) GetUser(ID int, args ...string) (user *User, err error) {
	var userID string
	if ID != 0 {
		userID = strconv.Itoa(ID)
	}

	var selections string
	for _, arg := range args {
		selections += arg + ","
	}
	selections = strings.TrimSuffix(selections, ",")

	data, err := s.callAPI(apiUser+endpoint(userID), map[string]string{"selections": selections})
	if err != nil {
		return
	}
	user = &User{}
	json.Unmarshal(data, user)
	return
}

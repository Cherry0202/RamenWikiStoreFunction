package structs

type ResGooglePlace struct {
	Results []struct {
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			LocationType string `json:"location_type"`
			Bounds       struct {
				Northeast struct {
					Lat int `json:"lat"`
					Lng int `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat int `json:"lat"`
					Lng int `json:"lng"`
				} `json:"southwest"`
			} `json:"bounds"`
			Viewport struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"viewport"`
			Types interface{} `json:"types"`
		} `json:"geometry"`
		Name             string   `json:"name"`
		Icon             string   `json:"icon"`
		PlaceID          string   `json:"place_id"`
		Rating           float64  `json:"rating"`
		UserRatingsTotal int      `json:"user_ratings_total"`
		Types            []string `json:"types"`
		OpeningHours     struct {
			OpenNow bool `json:"open_now"`
		} `json:"opening_hours,omitempty"`
		Photos []struct {
			PhotoReference   string   `json:"photo_reference"`
			Height           int      `json:"height"`
			Width            int      `json:"width"`
			HTMLAttributions []string `json:"html_attributions"`
		} `json:"photos"`
		PriceLevel     int    `json:"price_level,omitempty"`
		BusinessStatus string `json:"business_status"`
		ID             string `json:"id"`
	} `json:"Results"`
	HTMLAttributions []interface{} `json:"HTMLAttributions"`
	NextPageToken    string        `json:"NextPageToken"`
}

type Rework struct {
	HTMLAttributions []interface{} `json:"HTMLAttributions"`
	NextPageToken    string        `json:"NextPageToken"`
	Result           []Result      `json:"result"`
}

type Result struct {
	PlaceID          string       `json:"place_id"` // phone number取得時に必要
	Name             string       `json:"name"`
	FormattedAddress string       `json:"formatted_address"`
	OpeningHours     OpeningHours `json:"opening_hours"`
	Geometry         Geometry     `json:"geometry"`
	Photos           []Photos     `json:"photos"`
}

type OpeningHours struct {
	OpenNow bool `json:"open_now"`
}

type Geometry struct {
	Location Location `json:"location"`
}

type Photos struct {
	PhotoReference string `json:"photo_reference"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

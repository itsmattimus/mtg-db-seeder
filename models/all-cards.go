package models

import "encoding/json"

/*
	FROM THE SCRYFALL DOCS: A JSON file containing every card
	object on Scryfall in every language.
*/

type AllCards struct {
	Object          string          `json:"object" gorm:"type:varchar(10)"`
	Sid             string          `json:"id" gorm:"type:varchar(50)"`
	OracleId        string          `json:"oracle_id" gorm:"type:varchar(50)"`
	MultiverseIds   json.RawMessage `json:"multiverse_ids" gorm:"type:json"`
	MtgoId          int             `json:"mtgo_id" gorm:"type:int"`
	MtgoFoilId      int             `json:"mtgo_foil_id" gorm:"type:int"`
	TcgplayerId     int             `json:"tcgplayer_id" gorm:"type:int"`
	CardmarketId    int             `json:"cardmarket_id" gorm:"type:int"`
	Name            string          `json:"name" gorm:"type:varchar(200)"`
	Lang            string          `json:"lang" gorm:"type:varchar(4)"`
	ReleasedAt      string          `json:"released_at" gorm:"type:varchar(45)"`
	Uri             string          `json:"uri" gorm:"type:varchar(100)"`
	ScryfallUri     string          `json:"scryfall_uri" gorm:"type:varchar(500)"`
	Layout          string          `json:"layout" gorm:"type:varchar(45)"`
	HighresImage    bool            `json:"highres_image" gorm:"type:tinyint"`
	ImageStatus     string          `json:"image_status" gorm:"type:varchar(45)"`
	ImageUris       json.RawMessage `json:"image_uris" gorm:"type:json"`
	ManaCost        string          `json:"mana_cost" gorm:"type:varchar(100)"`
	Cmc             float64         `json:"cmc" gorm:"type:int"`
	TypeLine        string          `json:"type_line" gorm:"type:varchar(100)"`
	OracleText      string          `json:"oracle_text" gorm:"type:varchar(1000)"`
	Power           string          `json:"power" gorm:"type:varchar(10)"`
	Toughness       string          `json:"toughness" gorm:"type:varchar(10)"`
	Colors          json.RawMessage `json:"colors" gorm:"type:json"`
	ColorIdentity   json.RawMessage `json:"color_identity" gorm:"type:json"`
	Keywords        json.RawMessage `json:"keywords" gorm:"type:json"`
	Legalities      json.RawMessage `json:"legalities" gorm:"type:json"`
	Games           json.RawMessage `json:"games" gorm:"type:json"`
	Reserved        bool            `json:"reserved" gorm:"type:tinyint"`
	Foil            bool            `json:"foil" gorm:"type:tinyint"`
	Nonfoil         bool            `json:"nonfoil" gorm:"type:tinyint"`
	Finishes        json.RawMessage `json:"finishes" gorm:"type:json"`
	Oversized       bool            `json:"oversized" gorm:"type:tinyint"`
	Promo           bool            `json:"promo" gorm:"type:tinyint"`
	Reprint         bool            `json:"reprint" gorm:"type:tinyint"`
	Variation       bool            `json:"variation" gorm:"type:tinyint"`
	SetId           string          `json:"set_id" gorm:"type:varchar(50)"`
	Set             string          `json:"set" gorm:"type:varchar(5)"`
	SetName         string          `json:"set_name" gorm:"type:varchar(50)"`
	SetType         string          `json:"set_type" gorm:"type:varchar(50)"`
	SetUri          string          `json:"set_uri" gorm:"type:varchar(100)"`
	SetSearch_uri   string          `json:"set_search_uri" gorm:"type:varchar(100)"`
	ScryfallSetUri  string          `json:"scryfall_set_uri" gorm:"type:varchar(100)"`
	RulingsUri      string          `json:"rulings_uri" gorm:"type:varchar(100)"`
	PrintsSearchUri string          `json:"prints_search_uri" gorm:"type:varchar(200)"`
	CollectorNumber string          `json:"collector_number" gorm:"type:varchar(45)"`
	Digital         bool            `json:"digital" gorm:"type:tinyint"`
	Rarity          string          `json:"rarity" gorm:"type:varchar(8)"`
	FlavorText      string          `json:"flavor_text" gorm:"type:varchar(500)"`
	CardBackId      string          `json:"card_back_id" gorm:"type:varchar(50)"`
	Artist          string          `json:"artist" gorm:"type:varchar(50)"`
	ArtistIds       json.RawMessage `json:"artist_ids" gorm:"type:json"`
	IllustrationId  string          `json:"illustration_id" gorm:"type:varchar(50)"`
	BorderColor     string          `json:"border_color" gorm:"type:varchar(20)"`
	Frame           string          `json:"frame" gorm:"type:varchar(45)"`
	FullArt         bool            `json:"full_art" gorm:"type:tinyint"`
	Textless        bool            `json:"textless" gorm:"type:tinyint"`
	Booster         bool            `json:"booster" gorm:"type:tinyint"`
	StorySpotlight  bool            `json:"story_spotlight" gorm:"type:tinyint"`
	EdhrecRank      int             `json:"edhrec_rank" gorm:"type:int"`
	Prices          json.RawMessage `json:"prices" gorm:"type:json"`
	RelatedUris     json.RawMessage `json:"related_uris" gorm:"type:json"`
}

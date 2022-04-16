package models

/*
	FROM THE SCRYFALL DOCS: A JSON file of Scryfall card objects that together
	contain all unique artworks. The chosen cards promote the best image scans.
*/

type Rulings struct {
	Object      string `json:"object" gorm:"type:varchar(10)"`
	OracleId    string `json:"oracle_id" gorm:"type:varchar(50)"`
	Source      string `json:"source" gorm:"type:varchar(10)"`
	PublishedAt string `json:"published_at" gorm:"type:varchar(10)"`
	Comment     string `json:"comment" gorm:"type:varchar(5000)"`
}

package entity

type Plan struct {
	ID     int64  `gorm:"primary_key;auto_increment" json:"id"`
	Campus string `gorm:"type:text" json:"campus"`
	Corpus string `gorm:"type:text" json:"corpus"`
	Floor  string `gorm:"type:text" json:"floor"`
	SVG    string `gorm:"type:text" json:"svg"`
	Graph  string `gorm:"type:text" json:"graph"`
}

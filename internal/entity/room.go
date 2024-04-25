package entity

type Room struct {
	ID          int64  `gorm:"primary_key;auto_increment" json:"id"`
	RUName      string `gorm:"size:255;not null" json:"ru_name"`
	ENName      string `gorm:"size:255;not null" json:"en_name"`
	Description string `gorm:"size:255;" json:"description"`
	RUType      string `gorm:"size:255;" json:"ru_type"`
	ENType      string `gorm:"size:255;" json:"en_type"`
	EntryID     string `gorm:"primary_key;auto_increment" json:"entry_id"`
}

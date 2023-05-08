package models

type City struct {
	ID        int     `json:"id" gorm:"primary_key:auto_increment"`
	Name      string  `json:"name" gorm:"type: varchar(40)"`
	StationID int     `json:"station_id"`
	Station   Station `json:"station"`
}

package entity

// Feature : config table for backend
type Feature struct {
	ID      int    `gorm:"AUTO_INCREMENT;primary_key" json:"id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

// TableName : mapping between entity and physical table
func (Feature) TableName() string {
	return "rc_config"
}

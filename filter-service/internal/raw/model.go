package raw

import "gorm.io/gorm"

type RawData struct {
	gorm.Model
	Valid string
	Data  string
}

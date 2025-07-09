package input

import "gorm.io/gorm"

type Fservise struct {
	gorm.Model
	Data string
}

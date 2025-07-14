package processor

import "gorm.io/gorm"

type ProcessResult struct {
	gorm.Model
	Data string
}

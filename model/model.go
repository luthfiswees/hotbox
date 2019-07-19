package model

import (
	"github.com/jinzhu/gorm"
)

type ReportEntry struct {
	gorm.Model
	ServiceName string
	ReportFile string
	Status string
	Timestamp string
}
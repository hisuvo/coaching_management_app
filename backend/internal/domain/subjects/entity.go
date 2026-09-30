package subjects

import "gorm.io/gorm"

type Subject struct {
	gorm.Model

	Name        string `gorm:"size:100;not null"`
	Code   		string `gorm:"size:50;not null"`
	Description *string `gorm:"type:text"`
}

// composite unique approach
// uniqueIndex:idx_coaching_subject_code 

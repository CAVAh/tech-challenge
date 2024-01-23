package models

import (
	"github.com/CAVAh/api-tech-challenge/src/core/domain/entities"
	"gorm.io/gorm"
)

type ProductCategory struct {
	gorm.Model
	Description string
}

func (pc ProductCategory) ToDomain() entities.ProductCategory {
	return entities.ProductCategory{
		ID:          int(pc.ID),
		Description: pc.Description,
	}
}

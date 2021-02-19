package repository

import (
	"shopee-tracker/config"
	"shopee-tracker/models"

	"gorm.io/gorm"
)

var categoryRepoInstance *categoryRepo

// CategoryRepo ...
type CategoryRepo interface {
	Save(category models.Category) error
}

type categoryRepo struct {
	db *gorm.DB
}

func (c *categoryRepo) Save(category models.Category) error {
	if err := c.db.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

// GetCategoryRepo ...
func GetCategoryRepo() CategoryRepo {
	if categoryRepoInstance == nil {
		categoryRepoInstance = &categoryRepo{
			config.GetDbConnection(),
		}
	}
	return categoryRepoInstance
}

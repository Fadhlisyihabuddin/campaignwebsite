package campaign

import "gorm.io/gorm"

type Repository interface {
	FindByID(id int) ([]Campaign, error)
	FindAll() ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByID(id int) ([]Campaign, error) {
	var campaign []Campaign

	err := r.db.Where("user_id=?", id).Preload("CampaignImages", "campaign_images.is_primary = true").Find(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

//db.Where("state = ?", "active").Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
//db.Where("state = ?", "active").Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
// SELECT * FROM users WHERE state = 'active';
// SELECT * FROM orders WHERE user_id IN (1,2) AND state NOT IN ('cancelled');

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = true").Find(&campaigns).Error
	// err := r.db.Where("is_primary = ?", "true").Preload("CampaignImages").Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

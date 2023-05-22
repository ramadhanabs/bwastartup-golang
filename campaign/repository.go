package campaign

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Campaign, error)
	GetByUserId(UserID int) ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Campaign, error) {
	campaigns := []Campaign{}

	err := r.db.Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) GetByUserId(UserID int) ([]Campaign, error) {
	campaigns := []Campaign{}

	err := r.db.Where("user_id = ?", UserID).Preload("CampaignImages", "campaign_images.is_primary = TRUE").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

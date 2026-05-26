package models

import "gorm.io/gorm"

func AddWeightEntry(db *gorm.DB, userID uint, newWeight float64) error {
	entry := WeightEntry{Weight: newWeight, UserID: userID}
	if err := db.Create(&entry).Error; err != nil {
		return err
	}

	return db.Model(&User{}).Where("id = ?", userID).Update("weight", newWeight).Error
}

func GetWeightHistory(db *gorm.DB, userID uint) ([]WeightEntry, error) {
	var history []WeightEntry
	err := db.Where("user_id = ?", userID).Order("created_at desc").Find(&history).Error
	return history, err
}

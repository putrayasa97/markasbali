package model

import "gorm.io/gorm"

type Car struct {
	Model
	Nama  string `gorm:"not null" json:"nama"`
	Tipe  string `gorm:"not null" json:"tipe"`
	Tahun string `gorm:"not null" json:"tahun"`
}

func (cr *Car) Create(db *gorm.DB) error {
	err := db.Model(Car{}).Create(&cr).Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *Car) GetByID(db *gorm.DB) (Car, error) {
	car := Car{}
	err := db.Model(Car{}).Where("id = ?", cr.Model.ID).Take(&car).Error

	if err != nil {
		return Car{}, err
	}

	return car, nil
}

func (cr *Car) GetAll(db *gorm.DB) ([]Car, error) {
	cars := []Car{}

	err := db.Model(Car{}).Find(&cars).Error

	if err != nil {
		return []Car{}, err
	}

	return cars, nil
}

func (cr *Car) UpdateOneByID(db *gorm.DB) error {
	err := db.Model(Car{}).
		Select(
			"nama",
			"tipe",
			"tahun").
		Where("id = ?", cr.Model.ID).
		Updates(map[string]interface{}{
			"nama":  cr.Nama,
			"tipe":  cr.Tipe,
			"tahun": cr.Tahun,
		}).Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *Car) DeleteByID(db *gorm.DB) error {
	err := db.Model(Car{}).
		Where("id = ?", cr.Model.ID).
		Delete(&cr).Error

	if err != nil {
		return err
	}

	return nil
}

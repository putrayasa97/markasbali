package model

import "gorm.io/gorm"

type Car struct {
	Model
	BrandID      uint
	Nama         string `gorm:"not null" json:"nama"`
	Tipe         string `gorm:"not null" json:"tipe"`
	Tahun        string `gorm:"not null" json:"tahun"`
	Color        string `gorm:"not null" json:"color"`
	Condition    string `gorm:"not null" json:"condition"`
	UUID         string `gorm:"not null" json:"uuid"`
	SellingPrice int    `gorm:"not null; default:0" json:"selling_price"`
}

func (cr *Car) Create(db *gorm.DB) error {
	err := db.
		Model(Car{}).
		Create(&cr).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *Car) GetByUUID(db *gorm.DB) (Car, error) {
	res := Car{}

	err := db.
		Model(Car{}).
		Where("uuid = ?", cr.UUID).
		Take(&res).
		Error

	if err != nil {
		return Car{}, err
	}

	return res, nil
}

func (cr *Car) GetByID(db *gorm.DB) (Car, error) {
	res := Car{}

	err := db.
		Model(Car{}).
		Where("id = ?", cr.Model.ID).
		Take(&res).
		Error

	if err != nil {
		return Car{}, err
	}

	return res, nil
}

func (cr *Car) GetAll(db *gorm.DB) ([]Car, error) {
	res := []Car{}

	err := db.
		Model(Car{}).
		Find(&res).
		Error

	if err != nil {
		return []Car{}, err
	}

	return res, nil
}

func (cr *Car) UpdateOneByID(db *gorm.DB) error {
	err := db.
		Model(Car{}).
		Select("nama", "tipe", "tahun", "color", "condition", "uuid", "selling_price").
		Where("id = ?", cr.Model.ID).
		Updates(map[string]any{
			"uuid":          cr.UUID,
			"nama":          cr.Nama,
			"tipe":          cr.Tipe,
			"tahun":         cr.Tahun,
			"color":         cr.Color,
			"condition":     cr.Condition,
			"selling_price": cr.SellingPrice,
		}).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *Car) DeleteByID(db *gorm.DB) error {
	err := db.
		Model(Car{}).
		Where("id = ?", cr.Model.ID).
		Delete(&cr).
		Error

	if err != nil {
		return err
	}

	return nil
}

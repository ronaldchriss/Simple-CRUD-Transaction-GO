package item

import (
	"gorm.io/gorm"
)

type Reprository interface {
	Save(item Item) (Item, error)
	FindAll() ([]Item, error)
	FindByID(ID int) (Item, error)
	Delete(ID int) (Item, error)
	Update(item Item) (Item, error)
}

type reprository struct {
	db *gorm.DB
}

func NewReprository(db *gorm.DB) *reprository {
	return &reprository{db}
}

func (r *reprository) FindAll() ([]Item, error) {
	var items []Item
	err := r.db.Find(&items).Error
	if err != nil {
		return items, err
	}
	return items, nil
}

func (r *reprository) Save(item Item) (Item, error) {
	err := r.db.Create(&item).Error
	if err != nil {
		return item, err
	}
	return item, nil
}

func (r *reprository) FindByID(ID int) (Item, error) {
	var item Item

	err := r.db.Where("id=?", ID).Find(&item).Error

	if err != nil {
		return item, err
	}

	return item, nil
}

func (r *reprository) Delete(ID int) (Item, error) {
	var item Item

	err := r.db.Where("id=?", ID).Delete(&item).Error

	if err != nil {
		return item, err
	}

	return item, nil
}

func (r *reprository) Update(item Item) (Item, error) {
	err := r.db.Save(&item).Error
	if err != nil {
		return item, err
	}
	return item, nil
}

package models

import (
	"github.com/go-playground/validator"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	Item struct {
		ID          uint    `json:"id" gorm:"primaryKey,autoIncrement"`
		ImageURL    string  `json:"url"`
		Title       string  `json:"title" validate:"required" gorm:"unique;not null"`
		Description string  `json:"description"`
		Price       float64 `json:"price" validate:"required,gte=0" gorm:"not null"`
	}

	ItemRepo struct {
		DB *gorm.DB
	}
)

func NewItemRepo(dsn string) *ItemRepo {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Item{})
	return &ItemRepo{
		DB: db,
	}
}

func (ir *ItemRepo) CreateItem(i *Item) error {
	result := ir.DB.Create(i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ir *ItemRepo) GetAllItems() []Item {
	items := []Item{}
	ir.DB.Find(&items)
	return items

}

func (ir *ItemRepo) GetItemById(id uint) Item {
	var i Item
	ir.DB.Where("id = ?", id).First(&i)
	return i
}

func (ir *ItemRepo) UpdateItem(id uint, updateItem *Item) error {
	var i Item
	ir.DB.Where("id = ?", id).First(&i)
	updateItem.ID = i.ID
	i.Title = updateItem.Title
	i.Description = updateItem.Description
	i.Price = updateItem.Price
	result := ir.DB.Save(&i)
	return result.Error

}

func (ir *ItemRepo) DeleteItem(id uint) error {
	var i Item
	ir.DB.Where("id = ?", id).First(&i)
	result := ir.DB.Delete(&i)
	return result.Error
}

func (i *Item) Validate() error {
	validator := validator.New()
	return validator.Struct(i)
}

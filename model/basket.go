package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Basket struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Data      []byte    `json:"data,omitempty"`
	State     string    `json:"state"`
}

func GetAllBasket(db *gorm.DB) (*[]Basket, error) {
	var baskets []Basket
	var result = db.Find(&baskets)
	if result.Error != nil {
		return nil, errors.New("error retrieving baskets")
	}
	return &baskets, nil
}

func CreateBasket(db *gorm.DB, basket *Basket) error {
	var result = db.Create(basket)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateBasket(db *gorm.DB, basketID int64, updatedBasket *Basket) error {
	var existingBasket Basket
	var result = db.First(&existingBasket, basketID)
	if result.Error != nil {
		return errors.New("Basket not found")
	}
	// Update fields
	existingBasket.State = updatedBasket.State
	// Update other fields as needed
	result = db.Save(&existingBasket)
	if result.Error != nil {
		return errors.New("error updating basket")
	}
	return nil
}

func GetBasket(db *gorm.DB, basketID int64) (*Basket, error) {
	var basket Basket
	var result = db.First(&basket, basketID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &basket, nil
}

func DeleteBasket(db *gorm.DB, basketID int64) error {
	var basket Basket
	var result = db.First(&basket, basketID)
	if result.Error != nil {
		return errors.New("Basket not found")
	}
	result = db.Delete(&basket)
	if result.Error != nil {
		return errors.New("error deleting basket")
	}
	return nil
}

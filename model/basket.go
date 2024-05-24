package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type BasketState string

const (
	PENDING   BasketState = "PENDING"
	COMPLETED BasketState = "COMPLETED"
)

func isValidState(state BasketState) error {
	switch state {
	case PENDING, COMPLETED:
		return nil
	default:
		return BasketInvalidDataError{errors.New("invalid data")}
	}
}

type Basket struct {
	ID        int64       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time   `json:"created_at"` // gorm:autoCreateTime
	UpdatedAt time.Time   `json:"updated_at"` // gorm:autoUpdateTime
	Data      []byte      `json:"data,omitempty" gorm:"size:2048"`
	State     BasketState `json:"state" gorm:"default:PENDING"`
}

func GetAllBasket(db *gorm.DB) (*[]Basket, error) {
	var baskets []Basket
	var result = db.Find(&baskets)
	if result.Error != nil {
		return nil, BasketRetrieveError{errors.New("error retrieving baskets")}
	}
	return &baskets, nil
}

func CreateBasket(db *gorm.DB, basket *Basket) (*Basket, error) {
	var verr = isValidState(basket.State)
	if verr != nil {
		return nil, verr
	}
	var result = db.Create(basket)
	if result.Error != nil {
		return nil, result.Error
	}
	return basket, nil
}

func UpdateBasket(db *gorm.DB, basketID int64, updatedBasket *Basket) error {
	var existingBasket Basket
	var result = db.First(&existingBasket, basketID)
	if result.Error != nil {
		return BasketNotFoundError{errors.New("basket not found")}
	}
	var verr = isValidState(updatedBasket.State)
	if verr != nil {
		return verr
	}
	if existingBasket.State == COMPLETED {
		return BasketCompletedError{errors.New("basket is completed")}
	}
	existingBasket.State = updatedBasket.State
	existingBasket.Data = updatedBasket.Data
	// updated_at will be handled by gorm
	result = db.Save(&existingBasket)
	if result.Error != nil {
		return BasketUpdateError{errors.New("error updating basket")}
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
		return BasketNotFoundError{errors.New("basket not found")}
	}
	result = db.Delete(&basket)
	if result.Error != nil {
		return BasketDeleteError{errors.New("error deleting basket")}
	}
	return nil
}

package basket

import "time"

type Basket struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Data      []byte    `json:"data,omitempty"`
	State     string    `json:"state"`
}

func GetAllBasket() {

}

func CreateBasket() {

}

func UpdateBasket() {

}

func GetBasket() {

}

func DeleteBasket() {

}

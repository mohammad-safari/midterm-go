package main

import (
	"basket-keeper/web"
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var _, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var e = echo.New()
	e.GET("/basket", baskethandlers.GetAllBasket)
	e.POST("/basket", baskethandlers.CreateBasket)
	e.PATCH("/basket/:id", baskethandlers.UpdateBasket)
	e.GET("/basket/:id", baskethandlers.GetBasket)
	e.DELETE("/basket/:id", baskethandlers.DeleteBasket)
	e.Logger.Fatal(e.Start(":8080"))
}

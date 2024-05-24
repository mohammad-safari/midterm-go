package main

import (
	"basket-keeper/model"
	"basket-keeper/util"
	"basket-keeper/web"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	var db, err = util.ConnectToSQLite()
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	err = db.AutoMigrate(&model.Basket{})
	if err != nil {
		log.Fatal(err)
	}

	var e = echo.New()
	e.GET("/basket", web.GetAllBasket)
	e.POST("/basket", web.CreateBasket)
	e.PATCH("/basket/:id", web.UpdateBasket)
	e.GET("/basket/:id", web.GetBasket)
	e.DELETE("/basket/:id", web.DeleteBasket)
	e.Logger.Fatal(e.Start(":8080"))
}

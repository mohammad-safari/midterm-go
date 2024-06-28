package main

import (
	"basket-keeper/model"
	"basket-keeper/util"
	"basket-keeper/web"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	var db, err = util.ConnectToSQLite()
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	err = db.AutoMigrate(&model.Basket{}, &model.User{})
	if err != nil {
		log.Fatal(err)
	}

	var e = echo.New()
	e.GET("/basket", web.JwtMiddleware(web.GetAllBasket, false))
	e.POST("/basket", web.JwtMiddleware(web.CreateBasket, false))
	e.PATCH("/basket/:id", web.JwtMiddleware(web.UpdateBasket, false))
	e.GET("/basket/:id", web.JwtMiddleware(web.GetBasket, false))
	e.DELETE("/basket/:id", web.JwtMiddleware(web.DeleteBasket, false))
	e.GET("/user/auth", web.LoginUser)
	e.POST("/user", web.CreateUser)
	e.DELETE("/user", web.JwtMiddleware(web.DeleteUser, true))
	e.Validator = util.NewCustomValidator()
	e.Logger.Fatal(e.Start(":8080"))
}

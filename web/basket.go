package web

import (
	"basket-keeper/model"
	"basket-keeper/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllBasket(c echo.Context) error {
	var db, err = util.ConnectToSQLite()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error connecting to the database")
	}
	// defer db.Close()
	var baskets, gerr = model.GetAllBasket(db)
	if gerr != nil {
		return c.String(http.StatusInternalServerError, "Error retrieving baskets")
	}
	return c.JSON(http.StatusOK, baskets)
}

func CreateBasket(c echo.Context) error {
	var basket model.Basket
	if err := c.Bind(&basket); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request data")
	}
	var db, err = util.ConnectToSQLite()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error connecting to the database")
	}
	// defer db.Close()
	var created_basket, serr = model.CreateBasket(db, &basket)
	if serr.Error() == "invalid data" {
		return c.String(http.StatusBadRequest, "Invalid request data")
	}
	if serr != nil {
		return c.String(http.StatusInternalServerError, "Error creating basket")
	}
	return c.JSON(http.StatusCreated, created_basket)
}

func UpdateBasket(c echo.Context) error {
	var basketID, cerr = util.ConvertStr2Int(c.Param("id")) // Extract basket ID from the request
	if cerr != nil {
		return c.String(http.StatusBadRequest, "Invalid basket Id")
	}
	var updatedBasket model.Basket
	if err := c.Bind(&updatedBasket); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request data")
	}
	var db, err = util.ConnectToSQLite()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error connecting to the database")
	}
	// defer db.Close()
	var uerr = model.UpdateBasket(db, basketID, &updatedBasket)
	if uerr != nil {
		switch uerr.Error() {
		case "Basket not found":
			return c.String(http.StatusNotFound, "Basket not found")
		case "invalid data":
			return c.String(http.StatusBadRequest, "Invalid request data")
		case "Basket is Completed":
			return c.String(http.StatusUnprocessableEntity, "Basket is Completed already")
		case "error updating basket":
			return c.String(http.StatusInternalServerError, "Error updating basket")
		}
	}
	return c.String(http.StatusOK, "Basket updated successfully")
}

func GetBasket(c echo.Context) error {
	var basketID, cerr = util.ConvertStr2Int(c.Param("id")) // Extract basket ID from the request
	if cerr != nil {
		return c.String(http.StatusBadRequest, "Invalid basket Id")
	}
	var db, err = util.ConnectToSQLite()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error connecting to the database")
	}
	// defer db.Close()
	var basket, gerr = model.GetBasket(db, basketID)
	if gerr != nil {
		return c.String(http.StatusNotFound, "Basket not found")
	}
	// Return basket details as a response (e.g., JSON or HTML)
	return c.JSON(http.StatusOK, basket)
}

func DeleteBasket(c echo.Context) error {
	var basketID, cerr = util.ConvertStr2Int(c.Param("id")) // Extract basket ID from the request
	if cerr != nil {
		return c.String(http.StatusBadRequest, "Invalid basket Id")
	}
	var db, err = util.ConnectToSQLite()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error connecting to the database")
	}
	// defer db.Close()
	var derr = model.DeleteBasket(db, basketID)
	if derr != nil {
		switch derr.Error() {
		case "Basket not found":
			return c.String(http.StatusNotFound, "Basket not found")
		case "error deleting basket":
			return c.String(http.StatusInternalServerError, "Error deleting basket")
		}
	}
	return c.String(http.StatusOK, "Basket deleted successfully")
}

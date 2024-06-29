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
	var userId, _ = GetUserIDFromContext(c)
	var baskets, gerr = model.GetAllBasket(db, userId)
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
	var userId, _ = GetUserIDFromContext(c)
	var created_basket, serr = model.CreateBasket(db, userId, &basket)
	if serr != nil {
		switch serr.(type) {
		case model.BasketInvalidDataError:
			return c.String(http.StatusBadRequest, "Invalid request data")
		}
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
	var userId, _ = GetUserIDFromContext(c)
	var uerr = model.UpdateBasket(db, userId, basketID, &updatedBasket)
	if uerr != nil {
		switch uerr.(type) {
		case model.BasketNotFoundError:
			return c.String(http.StatusNotFound, "Basket not found")
		case model.BasketInvalidDataError:
			return c.String(http.StatusBadRequest, "Invalid request data")
		case model.BasketCompletedError:
			return c.String(http.StatusUnprocessableEntity, "Basket is Completed already")
		case model.BasketUpdateError:
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
	var userId, _ = GetUserIDFromContext(c)
	var basket, gerr = model.GetBasket(db, userId, basketID)
	if gerr != nil {
		return c.String(http.StatusNotFound, "Basket not found")
	}
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
	var userId, _ = GetUserIDFromContext(c)
	var derr = model.DeleteBasket(db, userId, basketID)
	if derr != nil {
		switch derr.(type) {
		case model.BasketNotFoundError:
			return c.String(http.StatusNotFound, "Basket not found")
		case model.BasketDeleteError:
			return c.String(http.StatusInternalServerError, "Error deleting basket")
		}
	}
	return c.String(http.StatusNoContent, "Basket deleted successfully")
}

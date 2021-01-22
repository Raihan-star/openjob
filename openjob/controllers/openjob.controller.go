package controllers

import (
	"net/http"
	"openjob/models"
	"strconv"

	"github.com/labstack/echo"
)

func FetchAllOpenjob(c echo.Context) error {
	result, err := models.FetchAllOpenjob()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreOpenjob(c echo.Context) error {
	code := c.FormValue("code")
	nama := c.FormValue("nama")
	keterangan := c.FormValue("keterangan")
	status := c.FormValue("status")
	tanggal := c.FormValue("tanggal")

	result, err := models.StoreOpenjob(code, nama, keterangan, status, tanggal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func UpdateOpenjob(c echo.Context) error {
	code := c.FormValue("code")
	nama := c.FormValue("nama")
	keterangan := c.FormValue("keterangan")
	status := c.FormValue("status")
	tanggal := c.FormValue("tanggal")

	conv_code, err := strconv.Atoi(code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateOpenjob(conv_code, nama, keterangan, status, tanggal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func DeleteOpenjob(c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteOpenjob(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

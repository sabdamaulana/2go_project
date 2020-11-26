package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"project_2go/models"
	"strconv"
)

func FetchAllMahasiswa(c echo.Context) error {
	result, err := models.FetchAllMahasiswa()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreMahasiswa(c echo.Context) error {
	nama := c.FormValue("nama")
	nim := c.FormValue("nim")
	email := c.FormValue("email")

	result, err := models.StoreMahasiswa(nama, nim, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func UpdateMahasiswa(c echo.Context) error {
	id := c.FormValue("id")
	nama := c.FormValue("nama")
	nim := c.FormValue("nim")
	email := c.FormValue("email")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateMahasiswa(conv_id, nama, nim, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func DeleteMahasiswa (c echo.Context) error{
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteMahasiswa(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

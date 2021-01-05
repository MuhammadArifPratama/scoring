package controllers

import (
	"github.com/labstack/echo/v4"
	"encoding/json"
	"tester/models"
	"net/http"
	log "github.com/Sirupsen/logrus"
)

func PMKCheking(pmk echo.Context) error {
	decode := json.NewDecoder(pmk.Request().Body)
	var data models.PMK
	err := decode.Decode(data)
	if err !=nil {
		return pmk.JSON(http.StatusMethodNotAllowed,err)
		log.Println(err)
	}
	return pmk.JSON(http.StatusOK,data)
}

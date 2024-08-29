package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Cron struct {
	TickSeconds uint   `json:"tick_seconds"`
	Command     string `json:"command"`
}

func (s *CronServer) addCron(c echo.Context) error {
	cron := new(Cron)
	if err := c.Bind(cron); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, cron)
}

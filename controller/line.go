package controller

import (
	"influx-alert-api/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Line Test
// @Tags Alert
// @Accept  json
// @Produce  json
// @Param Body body map[string]interface{} true "receive json from influx"
// @Success 200 {object} string
// @Router /api/Alert/Line [post]
func AlertLine(c *gin.Context) {

	body := new(map[string]interface{})

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		log.Printf("[ERROR] AlertLine unable to parse body, err: %s\n", err.Error())
		return
	}

	// log.Println(*body)

	received := *body
	alertMsg := received["_message"]

	res := services.SentLineNotify(alertMsg.(string))

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}

	c.JSON(http.StatusOK, res.Msg)
}

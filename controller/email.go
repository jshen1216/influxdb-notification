package controller

import (
	"fmt"
	"influx-alert-api/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Email Test
// @Tags Alert
// @Accept  json
// @Produce  json
// @Param Body body map[string]interface{} true "receive json from influx"
// @Success 200 {object} string
// @Router /api/Alert/Email [post]
func AlertEmail(c *gin.Context) {

	body := new(map[string]interface{})

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		log.Printf("[ERROR] AlertEmail unable to parse body, err: %s\n", err.Error())
		return
	}

	// log.Println(*body)

	received := *body
	alertMsg := received["_message"].(string)
	alertSubject := fmt.Sprintf("[Alert] Rule: %s & Level: %s", received["_check_name"].(string), received["_level"].(string))

	res := services.SendEmail(alertSubject, alertMsg)

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}

	c.JSON(http.StatusOK, res.Msg)
}

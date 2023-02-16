package services

import (
	"encoding/json"
	"influx-alert-api/global"
	"influx-alert-api/models"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func SentLineNotify(message string) models.Response {

	res := models.Response{}
	res.Success = false

	body := strings.NewReader(url.Values{
		"message": []string{message},
	}.Encode())

	client := &http.Client{}

	//這邊可以任意變換 http method  GET、POST、PUT、DELETE
	req, err := http.NewRequest("POST", "https://notify-api.line.me/api/notify", body)
	if err != nil {
		log.Printf("[Error] New request failed, err: %s\n", err)
		res.Msg = err.Error()
		return res
	}
	req.Header.Add("Authorization", "Bearer "+global.EnvConfig.Line.Token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[Error] Request do failed, err: %s\n", err.Error())
		res.Msg = err.Error()
		return res
	}
	defer resp.Body.Close()

	sitemap, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[Error] Read response failed, err: %s\n", err)
		res.Msg = err.Error()
		return res
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(sitemap, &result)
	if err != nil {
		log.Printf("[Error] Parse response failed, err: %s\n", err)
		res.Msg = err.Error()
		return res
	}
	if result["status"].(float64) != 200 {
		log.Printf("[Error] Line API response error, err: %s\n", result["message"].(string))
		res.Msg = result["message"].(string)
		return res
	}

	res.Success = true
	res.Msg = "Success to send line notification"

	return res
}

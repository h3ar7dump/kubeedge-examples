package lamp_driver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const bearerToken = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJjYzEwN2FmMTZiM2U0MjM4YmMyNzZlMGQ3NGE4ZTIyMSIsImlhdCI6MTU5NzY1MjM4NywiZXhwIjoxOTEzMDEyMzg3fQ.NWUwCrLZQmLG0z0GT1BXddr8O1AsVoVsviQLSubuOiw"

func TurnON() {
	url := "http://192.168.31.193:8123/api/services/light/turn_on"
	method := "POST"

	payload := strings.NewReader("{\"entity_id\": \"light.xiaomi_bedside_lamp2\"}")

	client := &http.Client{}
	req, err := http.NewRequest(method, strings.TrimSpace(url), payload)

	if err != nil {
		fmt.Println("-------------------")
		fmt.Println(err)
		fmt.Println("-------------------")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearerToken)

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

func TurnOff() {
	url := "http://192.168.31.193:8123/api/services/light/turn_off"
	method := "POST"

	payload := strings.NewReader("{\"entity_id\": \"light.xiaomi_bedside_lamp2\"}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearerToken)

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

}

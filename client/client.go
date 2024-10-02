package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"encoding/json"

	"github.com/naufalkhairil/golang-oauth-keycloak/config"
)

func InitClient(email string) {
	tokenUrl := config.GetAuthTokenURL(email)

	// Generate token for new login or
	// extend existing user token expiration
	//
	// base url redirect to auth url
	status, body := HitAPI(config.GetBaseURL())
	if status != 200 {
		log.Fatalf("Failed to init client: %d", status)
	}

	fmt.Println(string(body))

	// Check if login success until timeout
	startTs := time.Now()
	for {
		status, body := HitAPI(tokenUrl)
		if status == 200 {
			log.Print("Client authenticated")
			fmt.Printf("\n%s", body)

			// responseMap, err := ParseResponse(body)
			// if err != nil {
			// 	log.Fatalf("Failed to parse token response: %s", err)
			// }

			// fmt.Println(responseMap)

			break
		}

		if time.Since(startTs) > config.GetAuthTimeout() {
			log.Fatal("Timeout, failed to authenticate client")
		}

		time.Sleep(1 * time.Second)
	}
}

func ParseResponse(body []byte) (map[string]interface{}, error) {

	var responseMap map[string]interface{}
	if err := json.Unmarshal(body, &responseMap); err != nil {
		return nil, err
	}

	return responseMap, nil
}

func HitAPI(url string) (int, []byte) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	return resp.StatusCode, body
}

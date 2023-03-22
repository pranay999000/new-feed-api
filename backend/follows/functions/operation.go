package functions

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/pranay999000/follows/models"
)

func BasicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func CheckVertesExists(user_id string, channel chan bool) {
	url := "http://localhost:2480/command/UserGraph/sql"
	method := "POST"

	user_byte := []byte(user_id)
	var reqBody = []byte(`{"command": "select from Follows where user_id = :user_id", "parameters": {"user_id": "`)
	reqBody = append(reqBody, user_byte...)
	var end = []byte(`",}}`)
	reqBody = append(reqBody, end...)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Authorization", "Basic " + BasicAuth("root", "password"))

	if err != nil {
		log.Fatalln(err)
		channel <- false
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
		channel <- false
	}
	defer res.Body.Close()
	
	var result models.User

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
		channel <- false
	}

    if err := json.Unmarshal(body, &result); err != nil {
        log.Fatalln(err)
		channel <- false
    }

	channel <- len(result.Result) != 0

}
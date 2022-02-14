package payone

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var (
	api_url string = "https://api.pay1.de/post-gateway/"
)

type Response struct {
	Status      string `json:"status"`
	TxId        string `json:"txid"`
	UserId      string `json:"userid"`
	RedirectUrl string `json:"redirecturl"`
	Token       string `json:"token"`
}

func sendRequest(parameters url.Values) (Response, error) {

	req, err := http.NewRequest(http.MethodPost, api_url, strings.NewReader(parameters.Encode()))
	if err != nil {
		return Response{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("accept", "application/json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return Response{}, err
		}
		var response Response
		errUnmarshal := json.Unmarshal([]byte(body), &response)
		if errUnmarshal != nil {
			log.Println(string(body))
			return Response{}, errUnmarshal
		}
		return response, nil
	}
	return Response{}, fmt.Errorf("StatusCode is %d", resp.StatusCode)
}

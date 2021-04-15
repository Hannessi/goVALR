package goVALRapi

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type method string

func (m method) String() string {
	return string(m)
}

const GET method = "GET"
const POST method = "POST"
const PATCH method = "PATCH"

func HttpRequestWrapper(method method, url string, requestBody interface{}, response interface{}, apiKey string, apiSecret string) error {
	var body io.Reader
	marshaledBody := ""

	if requestBody != nil {
		marshaledBody, err := json.Marshal(requestBody)
		if err != nil {
			return err
		}

		body = bytes.NewBuffer(marshaledBody)
	}

	httpRequest, err := http.NewRequest(method.String(), url, body)
	if err != nil {
		return errors.New("could not create new HTTP request: " + err.Error())
	}

	httpRequest.Header.Add("Content-type", "application/json")
	if apiKey != "" {
		base := "https://api.valr.com"
		// todo: find a more elegant way to do this
		pathForSignature := strings.Replace(url, base, "", 1)
		signature, timestamp := createSignature(apiSecret, time.Now(), method.String(), pathForSignature, marshaledBody)

		httpRequest.Header.Add("X-VALR-API-KEY", apiKey)
		httpRequest.Header.Add("X-VALR-SIGNATURE", signature)
		httpRequest.Header.Add("X-VALR-TIMESTAMP", timestamp)
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return errors.New("could not send HTTP request: " + err.Error())
	}

	defer httpResponse.Body.Close()
	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return errors.New("could not read body of response: " + err.Error())
	}

	// todo - would be ideal if empty response could be in the JSON structure of the object.
	if string(responseBody) == "" {
		return nil
	}
	//fmt.Println("Response: ",string(responseBody))
	err = json.Unmarshal(responseBody, response)
	if err != nil {
		return errors.New("could not unmarshal the body of the response: " + err.Error())
	}
	return nil
}

func createSignature(apiSecret string, timestamp time.Time, verb string, path string, body string) (string, string) {
	// Create a new Keyed-Hash Message Authentication Code (HMAC) using SHA512 and API Secret
	mac := hmac.New(sha512.New, []byte(apiSecret))
	// Convert timestamp to nanoseconds then divide by 1000000 to get the milliseconds
	timestampString := strconv.FormatInt(timestamp.UnixNano()/1000000, 10)

	mac.Write([]byte(timestampString))
	mac.Write([]byte(strings.ToUpper(verb)))
	mac.Write([]byte(path))
	mac.Write([]byte(body))
	// Gets the byte hash from HMAC and converts it into a hex string
	return hex.EncodeToString(mac.Sum(nil)), timestampString
}

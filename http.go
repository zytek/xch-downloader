package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func DoRequest(method string, endpoint *url.URL) ([]byte, error) {

	client := http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest(method, endpoint.String(), nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Set("User-Agent", "xch-downloader")
	res, getErr := client.Do(req)

	if getErr != nil {
		return nil, getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	if res.StatusCode != 200 {
		err := errors.New(fmt.Sprintf("bad status code: %v", res.StatusCode))
		return nil, err
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, err
	}

	return body, nil
}

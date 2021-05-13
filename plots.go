package main

import (
	"encoding/json"
	"log"
	"net/url"
	"path"
	"strconv"
)

type PlotStatus struct {
	Id int `json:"id"`
	Progress string `json:"progressPercentage"`
	Eta string `json:"eta"`
	Status string `json:"status"`
	Filename string `json:"filename"`
	AwaitingRemoval bool `json:"awaitingRemoval"`
}

type PlotList []PlotStatus

func getPlotStatuses(u, key string) PlotList {
	endpoint, _ := url.Parse(u)
	endpoint.Path = path.Join(endpoint.Path, "plots", key)

	body, err := DoRequest("GET", endpoint )

	if err != nil {
		log.Println("ERROR: could not get list of plots:", err)
		return nil
	}

	plots := PlotList{}
	jsonErr := json.Unmarshal(body, &plots)
	if jsonErr != nil {
		log.Println("Failed to parse response:", jsonErr)
		return nil
	}
	return plots
}

func UrlForPlotDownload(plot PlotStatus) *url.URL {
	plotURL, _ := url.Parse(config.URL)
	plotURL.Path = path.Join(plotURL.Path, "plots", strconv.Itoa(plot.Id), "download")
	return plotURL
}
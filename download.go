package main

import (
	"fmt"
	"github.com/cavaliercoder/grab"
	"log"
	"net/url"
	"path"
	"strconv"
	"sync"
	"time"
)

type Download struct {
	wg *sync.WaitGroup
	URL *url.URL
	PlotID int
	Filename string
	Request *grab.Request
	Response *grab.Response
	OutputDir string
}

func (d *Download) Start() {
	if d.URL != nil {
		log.Printf("[%v] starting (%v/%v}", d.PlotID, d.OutputDir,d.Filename)
		client := grab.NewClient()
		req, err := grab.NewRequest(d.OutputDir, d.URL.String())
		if err != nil {
			log.Println(err)
		}
		d.Request = req
		d.Response = client.Do(req)
		go func() {
			<-d.Response.Done
			d.wg.Done()
		}()
	}
}

func (d *Download) Status() string {
	if d.Response != nil {
		return fmt.Sprintf("[%v] transferred %v / %v bytes (%.2f%%) %v/s, ETA %s",
			d.PlotID,
			bytes(d.Response.BytesComplete()),
			bytes(d.Response.Size),
			100*d.Response.Progress(),
			bytes(int64(d.Response.BytesPerSecond())),
			d.Response.ETA().Sub(time.Now()).Round(time.Second))
	} else {
		return fmt.Sprintf("[%v] pending..", d.PlotID)
	}
}

func (d *Download) MarkDone() {
	endpoint, _ := url.Parse(config.URL)
	endpoint.Path  = path.Join(endpoint.Path, "plots", strconv.Itoa(d.PlotID), config.FarmerKey)
	_, err := DoRequest("DELETE", endpoint)

	if err != nil {
		log.Println("ERROR: could not send DELETE request:", err)
	}
}

func Status(Downloads []*Download) {
	for _, v := range Downloads {
		if err := v.Response.Err(); err != nil {
			log.Printf("[%v] error during download: %v", err)
		} else {
			if config.Delete {
				log.Printf("[%v] plot downloaded, marking it for deletion")
				v.MarkDone()
			} else {
				log.Printf("[%v] plot downloaded, NOT deleting it on server due to -delete=false")
			}
		}
	}
}

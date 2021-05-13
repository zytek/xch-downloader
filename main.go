package main

import (
	"log"
	"sync"
	"time"
)
var config Config



func main() {
	config = parseFlags()

	t := time.NewTicker(1000 * time.Millisecond)
	defer t.Stop()

	for {
		log.Println("Checking plotting status")
		plots := getPlotStatuses(config.URL, config.FarmerKey)
		Downloads := make([]*Download, 0)
		wg := sync.WaitGroup{}

		for _, plot := range plots {
			if plot.Progress == "100" && !plot.AwaitingRemoval {
				log.Println("Plot [", plot.Id, "] is ready:", plot.Filename)
				d := Download{
					URL:      UrlForPlotDownload(plot),
					Filename: plot.Filename,
					PlotID:   plot.Id,
					OutputDir: config.DestinationDir,
					wg: &wg,
				}
				Downloads = append(Downloads, &d)
			}
		}

		log.Println("Downloading", len(Downloads), "files")
		for _, d := range Downloads {
			wg.Add(1)
			d.Start()
		}

		// start progress loop
		done := make(chan struct{})
		go func() {
			for {
				select {
				case <-t.C:
					for _, v := range Downloads {
						log.Println(v.Status())
					}
				case <-done:
					return
				}
			}
		}()
		// wait till all downloads finish
		wg.Wait()
		// stop progress loop
		done <- struct{}{}

		Status(Downloads)

		log.Println("All done, re-checking for new plots")
	}
}

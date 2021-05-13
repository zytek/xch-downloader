package main

import (
	"flag"
	"log"
	"os"
)

type Config struct {
	URL      string
	FarmerKey string
	DestinationDir string
	Delete bool
	ParallelDownloads    int
	Daemonize bool
}

func parseFlags() Config {

	config := Config{}
	flag.StringVar(&config.URL, "url", "https://xch-plotter.com", "plotter url")
	flag.BoolVar(&config.Delete, "delete", false, "send DELETE requests after downloading plot")
	flag.BoolVar(&config.Delete, "d", false, "daemonize (check and download continuously")
	flag.StringVar(&config.FarmerKey, "key", "", "farmer key")
	flag.StringVar(&config.DestinationDir, "dir", ".", "plot directory")
	flag.IntVar(&config.ParallelDownloads, "n", 10, "download n files simultaneously")
	flag.Parse()

	if len(config.FarmerKey)  != 96 {
		log.Println("ERROR: -key is required and expected to be 96 char long")
		flag.Usage()
		os.Exit(1)
	}

	return config
}

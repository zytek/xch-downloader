package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
)

func handleInterrupt(cancelFunc context.CancelFunc) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		log.Println("Received an interrupt, canceling copy...")
		log.Println("")
		cancelFunc()
		log.Println("Exiting")
		os.Exit(0)
	}()
}

func bytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

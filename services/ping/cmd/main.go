package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	url, exists := os.LookupEnv("URL")
	if !exists {
		log.Fatalln("URL environment variable not set")
	}

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Process canceled.")
			return
		case <-quit:
			fmt.Println("Received interrupt signal. Exiting...")
			cancel()
		default:
			resp, err := client.Get(url)
			if err != nil {
				log.Println("Ping failed:", err)
			} else {
				fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode), url)
				if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
					fmt.Println("HTTP Status is in the 2xx range")
					defer resp.Body.Close()
					bodyBytes, err := io.ReadAll(resp.Body)
					if err != nil {
						log.Println("Error reading response body:", err)
					} else {
						bodyString := string(bodyBytes)
						fmt.Println("Response Body:", bodyString)
					}
				} else {
					fmt.Printf(" %s Broken\n", url)
				}
				resp.Body.Close()
			}
			time.Sleep(5 * time.Second)
		}
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Iran struct {
	City       string `json:"city"`
	Population int64  `json:"population"`
	Men        int64  `json:"men"`
	Women      int64  `json:"women"`
	HOffset    int64  `json:"hOffset"`
	VOffset    int    `json:"vOffset"`
	Weather    string `json:"weather"`
	Timestamp  string `json:"timestamp"`
}

func main() {
	// Set up a channel to listen for interrupt signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		// Seed random number generation
		rand.Seed(time.Now().UnixNano())

		randomLongInt := rand.Int63n(1000000) + 1000000
		randomShortInt := rand.Intn(100) + 1

		men := randomLongInt / int64(randomShortInt)
		women := randomLongInt / 2

		// Get the current timestamp
		timestamp := time.Now().Format("2006-01-02 15:04:05")

		iran := Iran{
			City:       "tehran",
			Population: randomLongInt,
			Men:        men,
			Women:      women,
			HOffset:    randomLongInt,
			VOffset:    100,
			Weather:    "sun1.city = (100/90)*100",
			Timestamp:  timestamp,
		}

		jsonOutput, err := json.Marshal(map[string]interface{}{"iran": iran})
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			return
		}

		fmt.Println(string(jsonOutput))

		// Sleep for 1 second
		time.Sleep(1 * time.Second)

		// Check if an interrupt signal was received
		select {
		case <-sigChan:
			fmt.Println("\nInterrupt signal received. Exiting...")
			return
		default:
			// Continue if no interrupt signal
		}
	}
}

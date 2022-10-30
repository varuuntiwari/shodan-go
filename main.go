package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/varuuntiwari/shodan-go/scan"
)

func main() {
	var Iter int64
	var Filename string
	var wg sync.WaitGroup
	var i int64

	// Use random seed with UNIX time in nanoseconds
	rand.Seed(time.Now().UnixNano())

	// Get parameters
	flag.Int64Var(&Iter, "i", 10, "specify number of IPs to scan")
	flag.StringVar(&Filename, "f", "logs", "give filename to write logs to")
	flag.Parse()

	// Open file for logging
	w, err := os.OpenFile(Filename, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0644)
	if err != nil {
    	log.Fatalf("error opening log file: %v", err)
	}
	defer w.Close()
	log.SetOutput(w)

	// Use goroutines to optimize scans
	for i = 0; i < Iter; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr := fmt.Sprint(rand.Intn(254) + 1) + "." +
					fmt.Sprint(rand.Intn(254) + 1) + "." +
					fmt.Sprint(rand.Intn(254) + 1) + "." +
					fmt.Sprint(rand.Intn(254) + 1)
			scan.CheckIP(addr)
		}()
	}
	wg.Wait()
}
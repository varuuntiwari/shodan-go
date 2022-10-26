package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-ping/ping"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		addr := fmt.Sprint(rand.Intn(254) + 1) + "." + fmt.Sprint(rand.Intn(254) + 1) + "." + fmt.Sprint(rand.Intn(254) + 1) + "." + fmt.Sprint(rand.Intn(254) + 1)
		fmt.Print(addr, ": ")
		pinger, err := ping.NewPinger(addr)
		if err != nil {
			fmt.Println(err)
		} else {
			pinger.SetPrivileged(true)
			pinger.Count = 2
			pinger.Timeout = time.Second * 5
			err = pinger.Run()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%+v\n", pinger.Statistics())
			}
		}
	}
}
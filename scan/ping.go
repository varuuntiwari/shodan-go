package scan

import (
	"log"
	"time"

	"github.com/go-ping/ping"
)

// CheckIP checks whether an IP is reachable
// by sending ICMP packets.
func CheckIP(ip string) {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		log.Printf("[x] %v: %v\n", ip, err)
	} else {
		pinger.SetPrivileged(true)
		pinger.Count = 5
		pinger.Timeout = time.Second * 5
		err = pinger.Run()
		if err != nil {
			log.Printf("[x] %v: %v\n", ip, err)
		} else if(pinger.PacketsRecv > 0) {
			log.Printf("[+] %v: %+v\n", ip, pinger.Statistics())
		} else {
			log.Printf("[-] %v: %+v\n", ip, pinger.Statistics())
		}
	}
}
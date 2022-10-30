# shodan-go
 tool for randomly scanning and logging IPs, like shodan

 # Note
 Currently under development, it can only randomly generate IPs and ping them with 2 ICMP packets with a timeout of 5 seconds for each address. Runs on an endless for loop. 

 # Objectives
- [x] Log results to a file(s)
- [ ] Give the code some structure
- [ ] Avoid repeating IP addresses
- [x] Take parameters from command line
- [ ] Scan ports instead of pinging them
- [ ] Capable of scanning IPv6 addresses

# Dependencies
- [ping](https://github.com/go-ping/ping) package for pings.

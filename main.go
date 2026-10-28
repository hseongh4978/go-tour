package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ipAddr IPAddr) String() string {
	ip := make([]string, len(ipAddr))
	for i, _ip := range ipAddr {
		ip[i] = strconv.Itoa(int(_ip))
	}
	return strings.Join(ip, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

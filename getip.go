package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Failed to get nic info!")
		os.Exit(1)
	}
	for _, iface := range ifaces {
		fmt.Printf("HW: %-10s %v\n", iface.Name, iface.HardwareAddr)
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					fmt.Printf("    ADDR: %-10s %-15v %v\n", iface.Name, ipnet.IP, net.IP(ipnet.Mask))
				}
			}
		}
	}
}

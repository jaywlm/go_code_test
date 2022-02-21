package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ipv4Addr, ipv4Net, err := net.ParseCIDR("192.0.2.1/24")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv4Addr)
	fmt.Println(ipv4Net)
	fmt.Println("------")
	fmt.Printf("%T\n", ipv4Net.IP)
	fmt.Printf("%b\n", ipv4Net.Mask)
	fmt.Println("------")
	//fmt.Printf("%b\n", ipv4Net.IP & ipv4Net.Mask)
	//fmt.Printf("%b\n", ipv4Net.IP | ipv4Net.Mask)
}

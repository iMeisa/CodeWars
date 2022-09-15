package main

import (
	"CodeWars/countIPs"
	"fmt"
)

func main() {
	//fmt.Println(countIPs.GetIpVal("0.0.1.2"))
	fmt.Println(countIPs.IpsBetween("10.0.0.0", "10.0.0.50"))
	fmt.Println(countIPs.IpsBetween("20.0.0.10", "20.0.1.0"))
}

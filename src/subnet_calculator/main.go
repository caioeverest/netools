package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"encoding/binary"
	"net"
	"math"
)

func NumOfSetBits(n int) int{
   count := 0
   for n !=0{
      count += n &1
      n >>= 1
   }
   return count
}

func backtoIP4(ipInt uint32) string {
	ipInt64 := int64(ipInt)
	// need to do two bit shifting and “0xff” masking
	b0 := strconv.FormatInt((ipInt64>>24)&0xff, 10)
	b1 := strconv.FormatInt((ipInt64>>16)&0xff, 10)
	b2 := strconv.FormatInt((ipInt64>>8)&0xff, 10)
	b3 := strconv.FormatInt((ipInt64 & 0xff), 10)
	return b0 + "." + b1 + "." + b2 + "." + b3
}

func main() {
	fmt.Println()
	fmt.Println("subnet_calculator - Provides simple IP subnet calculations")
	fmt.Println()
	if len(os.Args) < 3 || net.ParseIP(os.Args[1]) == nil || net.ParseIP(os.Args[2]) == nil {
		fmt.Println("USAGE: subnet_calculator IP MASK")
		fmt.Println()
		fmt.Println("where")
		fmt.Println()
		fmt.Println("IP:    IP Address in dotted quad notation")
		fmt.Println("MASK:  Subnet mask in dotted quad notation")
		fmt.Println()
	} else {
		fmt.Println("IP address: ",os.Args[1])
		ipInt := binary.BigEndian.Uint32(net.ParseIP(os.Args[1])[12:16])
		fmt.Println("Subnet mask: ",os.Args[2])
		fmt.Println()
		maskInt := binary.BigEndian.Uint32(net.ParseIP(os.Args[2])[12:16])
		maskArray := strings.Split(os.Args[2],".")
		netInt := ipInt & maskInt
		fmt.Println("Network address: ", backtoIP4(netInt))
		invertedMaskInt := binary.BigEndian.Uint32(net.ParseIP(maskArray[3] + "." + maskArray[2] + "." + maskArray[1] + "." + maskArray[0])[12:16])
		broadcastInt := ipInt | invertedMaskInt
		fmt.Println("Broadcast address: ", backtoIP4(broadcastInt))
		cidrInt := NumOfSetBits(int(maskInt))
		hostsFloat := math.Pow(2, float64(32 - cidrInt)) - 2;	
		fmt.Println("Number of valid hosts per subnet: ", hostsFloat)
		quad255Int := binary.BigEndian.Uint32(net.ParseIP("255.255.255.255")[12:16])
		wildcardInt := quad255Int - maskInt
		fmt.Println("Wildcard mask: ", backtoIP4(wildcardInt))
		fmt.Println("Number of mask bits in CIDR notation: ", cidrInt)
		fmt.Println()
	}
}


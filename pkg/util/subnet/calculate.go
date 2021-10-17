package subnet

import (
	"encoding/binary"
	"math"
	"net"
	"strconv"
)

func numOfSetBits(n int) int {
	count := 0
	for n != 0 {
		count += n & 1
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

func IsValidSubnetMask(n uint32) bool {
	bitCount := 0
	setBitCount := 0
	for bitCount < 32 {
		bit := n & 1
		if setBitCount >= 1 {
			if bit == 0 {
				return false
			}
		} else {
			if bit == 1 {
				setBitCount += 1
			}
		}
		n >>= 1
		bitCount += 1
	}
	return true
}

func Calculate(ipAddress, subnetMask string) (string, string, int, string, int) {
	ipInt := binary.BigEndian.Uint32(net.ParseIP(ipAddress)[12:16])
	maskInt := binary.BigEndian.Uint32(net.ParseIP(subnetMask)[12:16])
	if !IsValidSubnetMask(maskInt) {
		panic("Invalid Subnet Mask")
	}
	netInt := ipInt & maskInt

	networkAddress := backtoIP4(netInt)
	invertedMaskInt := ^maskInt
	broadcastInt := ipInt | invertedMaskInt

	broadcastAddress := backtoIP4(broadcastInt)
	cidrInt := numOfSetBits(int(maskInt))
	hostsFloat := math.Pow(2, float64(32-cidrInt)) - 2

	noOfHosts := int(hostsFloat)
	quad255Int := binary.BigEndian.Uint32(net.ParseIP("255.255.255.255")[12:16])
	wildcardInt := quad255Int - maskInt

	wildcardMask := backtoIP4(wildcardInt)
	cidrNotation := cidrInt

	return networkAddress, broadcastAddress, noOfHosts, wildcardMask, cidrNotation
}

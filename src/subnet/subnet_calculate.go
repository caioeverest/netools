package subnet

import (
	"encoding/binary"
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

const (
	cliSubnetCalculate = "calculate"

	cliInvalidArgsErr = "invalid arguments provided for subnet calculate cmd"
)

func NumOfSetBits(n int) int {
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

func subNetCalculator(ipAddress, subnetMask string) (string, string, float64, string, int) {
	ipInt := binary.BigEndian.Uint32(net.ParseIP(ipAddress)[12:16])
	maskInt := binary.BigEndian.Uint32(net.ParseIP(subnetMask)[12:16])
	maskArray := strings.Split(subnetMask, ".")
	netInt := ipInt & maskInt

	networkAddress := backtoIP4(netInt)
	invertedMaskInt := binary.BigEndian.Uint32(net.ParseIP(maskArray[3] + "." + maskArray[2] + "." + maskArray[1] + "." + maskArray[0])[12:16])
	broadcastInt := ipInt | invertedMaskInt

	broadcastAddress := backtoIP4(broadcastInt)
	cidrInt := NumOfSetBits(int(maskInt))
	hostsFloat := math.Pow(2, float64(32-cidrInt)) - 2

	noOfHosts := hostsFloat
	quad255Int := binary.BigEndian.Uint32(net.ParseIP("255.255.255.255")[12:16])
	wildcardInt := quad255Int - maskInt

	wildcardMask := backtoIP4(wildcardInt)
	cidrNotation := cidrInt

	return networkAddress, broadcastAddress, noOfHosts, wildcardMask, cidrNotation
}

var CalculateSubnet = &cobra.Command{
	Use:   cliSubnetCalculate,
	Short: "CLI for subnet calculation operations",
	Long:  "CLI for subnet calculation operations",
	Args:  cobra.MaximumNArgs(100),
	Run: func(cmd *cobra.Command, args []string) {
		ipAddress, err := cmd.Flags().GetString("ipAddress")
		if err != nil {
			fmt.Println(err)
		}

		subnetMask, err := cmd.Flags().GetString("subnetMask")
		if err != nil {
			fmt.Println(err)
			return
		}

		if net.ParseIP(ipAddress) == nil || net.ParseIP(subnetMask) == nil {
			fmt.Println(cliInvalidArgsErr)
			return
		}

		networkAddress, broadcastAddress, noOfHosts, wildcardMask, cidrNotation :=
			subNetCalculator(ipAddress, subnetMask)

		fmt.Println("IP address: ", ipAddress)
		fmt.Println("Subnet mask: ", subnetMask)
		fmt.Println("Network address: ", networkAddress)
		fmt.Println("Broadcast address: ", broadcastAddress)
		fmt.Println("Number of valid hosts per subnet: ", noOfHosts)
		fmt.Println("Wildcard mask: ", wildcardMask)
		fmt.Println("Number of mask bits in CIDR notation: ", cidrNotation)
	},
}

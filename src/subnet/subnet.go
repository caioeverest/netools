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
	cliCmd = "subnet"

	cliInvalidArgsErr = "invalid arguments provided for subnet cli"
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

var Cli = &cobra.Command{
	Use:   cliCmd,
	Short: "CLI for subnet operations",
	Long:  "CLI for subnet operations",
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

		fmt.Println("IP address: ", ipAddress)
		ipInt := binary.BigEndian.Uint32(net.ParseIP(ipAddress)[12:16])
		fmt.Println("Subnet mask: ", subnetMask)

		maskInt := binary.BigEndian.Uint32(net.ParseIP(subnetMask)[12:16])
		maskArray := strings.Split(subnetMask, ".")
		netInt := ipInt & maskInt

		fmt.Println("Network address: ", backtoIP4(netInt))
		invertedMaskInt := binary.BigEndian.Uint32(net.ParseIP(maskArray[3] + "." + maskArray[2] + "." + maskArray[1] + "." + maskArray[0])[12:16])
		broadcastInt := ipInt | invertedMaskInt

		fmt.Println("Broadcast address: ", backtoIP4(broadcastInt))
		cidrInt := NumOfSetBits(int(maskInt))
		hostsFloat := math.Pow(2, float64(32-cidrInt)) - 2

		fmt.Println("Number of valid hosts per subnet: ", hostsFloat)
		quad255Int := binary.BigEndian.Uint32(net.ParseIP("255.255.255.255")[12:16])
		wildcardInt := quad255Int - maskInt

		fmt.Println("Wildcard mask: ", backtoIP4(wildcardInt))
		fmt.Println("Number of mask bits in CIDR notation: ", cidrInt)
	},
}

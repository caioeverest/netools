package route

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"

	"netools/pkg/util/subnet"
)

const (
	cliInvalidArgsErr = "invalid arguments provided for subnet calculate cmd"
)

// CmdCalculateSubnet to represent command: ./netools subnet calculate -i <ip address> -s <subnet mask>"
var CmdCalculateSubnet = &cobra.Command{
	Use:   "calculate",
	Short: "CLI for subnet calculation operations",
	Long:  "CLI for subnet calculation operations",
	Run: func(cmd *cobra.Command, args []string) {
		ipAddress, err := cmd.Flags().GetString("ipAddress")
		if err != nil {
			fmt.Println(err)
			return
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
			subnet.Calculate(ipAddress, subnetMask)

		fmt.Println("IP address: ", ipAddress)
		fmt.Println("Subnet mask: ", subnetMask)
		fmt.Println("Network address: ", networkAddress)
		fmt.Println("Broadcast address: ", broadcastAddress)
		fmt.Println("Number of valid hosts per subnet: ", noOfHosts)
		fmt.Println("Wildcard mask: ", wildcardMask)
		fmt.Println("Number of mask bits in CIDR notation: ", cidrNotation)
	},
}

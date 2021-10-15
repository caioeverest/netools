package subnet

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"

	"netools/pkg/model/subnet"
)

const (
	cmdCalculate = "calculate"

	cliInvalidArgsErr = "invalid arguments provided for subnet calculate cmd"
)

var CalculateSubnet = &cobra.Command{
	Use:   cmdCalculate,
	Short: "CLI for subnet calculation operations",
	Long:  "CLI for subnet calculation operations",
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
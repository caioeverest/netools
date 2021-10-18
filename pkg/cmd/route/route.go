package route

import "github.com/spf13/cobra"

const (
	errInvalidArgIpAddress  = "invalid ip address is provided"
	errInvalidArgSubnetMask = "invalid subnet mask is provided"
)

func Initialize(cmd *cobra.Command) {
	// Level 1
	cmd.AddCommand(cmdSubnet)

	// Level 2
	cmdSubnet.AddCommand(cmdSubnetCalculate)

	// Level 2 arg flags
	cmdSubnetCalculate.PersistentFlags().StringP("ipAddress", "i", "", "ip address")
	cmdSubnetCalculate.PersistentFlags().StringP("subnetMask", "s", "", "subnet mask")

	// Level 2 arg flags marked as required
	_ = cmdSubnetCalculate.MarkPersistentFlagRequired("ipAddress")
	_ = cmdSubnetCalculate.MarkPersistentFlagRequired("subnetMask")
}

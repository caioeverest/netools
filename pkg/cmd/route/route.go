package route

import "github.com/spf13/cobra"

const (
	errInvalidArgIpAddress  = "invalid ip address is provided"
	errInvalidArgSubnetMask = "invalid subnet mask is provided"

	flagIpaddress  = "ipAddress"
	flagSubnetMask = "subnetMask"
)

func Initialize(cmd *cobra.Command) {
	// Level 1
	cmd.AddCommand(cmdSubnet)

	// Level 2
	cmdSubnet.AddCommand(cmdSubnetCalculate)

	// Level 2 arg flags
	cmdSubnetCalculate.PersistentFlags().StringP(flagIpaddress, "i", "", "ip address")
	cmdSubnetCalculate.PersistentFlags().StringP(flagSubnetMask, "s", "", "subnet mask")

	// Level 2 arg flags marked as required
	_ = cmdSubnetCalculate.MarkPersistentFlagRequired(flagIpaddress)
	_ = cmdSubnetCalculate.MarkPersistentFlagRequired(flagSubnetMask)
}

package route

import "github.com/spf13/cobra"

func Initialize(cmd *cobra.Command) {
	// Level 1
	cmd.AddCommand(CmdSubnet)

	// Level 2
	CmdSubnet.AddCommand(CmdCalculateSubnet)

	// Level 2 arg flags
	CmdCalculateSubnet.PersistentFlags().StringP("ipAddress", "i", "", "ip address")
	CmdCalculateSubnet.PersistentFlags().StringP("subnetMask", "s", "", "subnet mask")

	// Level 2 arg flags marked as required
	_ = CmdCalculateSubnet.MarkPersistentFlagRequired("ipAddress")
	_ = CmdCalculateSubnet.MarkPersistentFlagRequired("subnetMask")
}

package cli

import (
	"netools/src/subnet"
)

func init() {
	rootCmd.AddCommand(subnet.Cli)

	subnet.Cli.PersistentFlags().StringP("ipAddress", "i", "", "ip address")
	subnet.Cli.PersistentFlags().StringP("subnetMask", "s", "", "subnet mask")

	_ = subnet.Cli.MarkPersistentFlagRequired("ipAddress")
	_ = subnet.Cli.MarkPersistentFlagRequired("subnetMask")
}

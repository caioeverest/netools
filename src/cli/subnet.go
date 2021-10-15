package cli

import (
	"netools/src/subnet"
)

func init() {
	rootCmd.AddCommand(subnet.Subnet)
}

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"netools/src/example"
	"netools/src/subnet"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "netools",
	Short: "Netools CLI App",
	Long:  `Welcome to netools CLI`,
}

func init() {
	// CLI and sub CLI's
	rootCmd.AddCommand(example.CliExample)
	rootCmd.AddCommand(subnet.Subnet)

	example.CliExample.AddCommand(example.SubCliExampleWithArgs)

	subnet.Subnet.AddCommand(subnet.CalculateSubnet)

	// Cli flags
	example.SubCliExampleWithArgs.PersistentFlags().StringP("message", "m", "", "example of sub cli with args")

	subnet.CalculateSubnet.PersistentFlags().StringP("ipAddress", "i", "", "ip address")
	subnet.CalculateSubnet.PersistentFlags().StringP("subnetMask", "s", "", "subnet mask")

	// CLI flags marked as required
	_ = example.SubCliExampleWithArgs.MarkPersistentFlagRequired("message")

	_ = subnet.CalculateSubnet.MarkPersistentFlagRequired("ipAddress")
	_ = subnet.CalculateSubnet.MarkPersistentFlagRequired("subnetMask")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

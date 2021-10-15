package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"netools/src/example"
	"netools/src/subnet"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "netools",
	Short: "Netools CLI App",
	Long:  `Welcome to netools CLI`,
}

func init() {
	// Root Cobra CLI
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rootcli.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Example CLI
	rootCmd.AddCommand(example.CliExample)
	example.CliExample.AddCommand(example.SubCliExampleWithArgs)
	example.SubCliExampleWithArgs.PersistentFlags().StringP("message", "m", "", "example of sub cli with args")
	_ = example.SubCliExampleWithArgs.MarkPersistentFlagRequired("message")

	// Subnet CLI
	rootCmd.AddCommand(subnet.Subnet)
	subnet.Subnet.AddCommand(subnet.CalculateSubnet)
	subnet.CalculateSubnet.PersistentFlags().StringP("ipAddress", "i", "", "ip address")
	subnet.CalculateSubnet.PersistentFlags().StringP("subnetMask", "s", "", "subnet mask")
	_ = subnet.CalculateSubnet.MarkPersistentFlagRequired("ipAddress")
	_ = subnet.CalculateSubnet.MarkPersistentFlagRequired("subnetMask")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

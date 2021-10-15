package route

import "github.com/spf13/cobra"

// CmdSubnet to represent command: ./netools subnet
var CmdSubnet = &cobra.Command{
	Use:   "subnet",
	Short: "Subnet CLI utility",
	Long:  `CLI tools to perform subnet operations`,
}

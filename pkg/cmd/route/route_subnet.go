package route

import "github.com/spf13/cobra"

// cmdSubnet represents
// "subnet"
var cmdSubnet = &cobra.Command{
	Use:   "subnet",
	Short: "Subnet CLI utility",
	Long:  `CLI tools to perform subnet operations`,
}

package subnet

import "github.com/spf13/cobra"

const (
	cliSubnet = "subnet"
)

var Subnet = &cobra.Command{
	Use:   cliSubnet,
	Short: "Subnet CLI tool",
	Long:  `CLI tools to perform subnet operations`,
}

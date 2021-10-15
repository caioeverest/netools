package subnet

import "github.com/spf13/cobra"

const (
	cmdSubnet = "subnet"
)

var Subnet = &cobra.Command{
	Use:   cmdSubnet,
	Short: "Subnet CLI tool",
	Long:  `CLI tools to perform subnet operations`,
}

package subnet

import "github.com/spf13/cobra"

const (
	cmdSubnet = "subnet"
)

var Subnet = &cobra.Command{
	Use:   cmdSubnet,
	Short: "Subnet CLI util",
	Long:  `CLI tools to perform subnet operations`,
}

package cli

import "netools/src/subnet"

func init() {
	subnet.Subnet.AddCommand(subnet.CalculateSubnet)

	subnet.CalculateSubnet.PersistentFlags().StringP("ipAddress", "i", "", "ip address")
	subnet.CalculateSubnet.PersistentFlags().StringP("subnetMask", "s", "", "subnet mask")

	_ = subnet.CalculateSubnet.MarkPersistentFlagRequired("ipAddress")
	_ = subnet.CalculateSubnet.MarkPersistentFlagRequired("subnetMask")
}

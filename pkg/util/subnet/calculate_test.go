package subnet

import (
	"encoding/binary"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	var tests = []struct {
		IPAddress                string
		SubnetMask               string
		ExpectedNetworkAddress   string
		ExpectedBroadcastAddress string
		ExpectedNoOfHosts        int
		ExpectedWildcardMask     string
		ExpectedCidrNotation     int
	}{
		{"198.168.1.101", "255.255.255.128", "198.168.1.0", "198.168.1.127", 126, "0.0.0.127", 25},
		{"198.168.1.101", "255.255.192.0", "198.168.0.0", "198.168.63.255", 16382, "0.0.63.255", 18},
		{"198.168.1.101", "255.248.0.0", "198.168.0.0", "198.175.255.255", 524286, "0.7.255.255", 13},
		{"52.12.12.52", "128.0.0.0", "0.0.0.0", "127.255.255.255", 2147483646, "127.255.255.255", 1},
		{"52.12.12.52", "128.0.0.0", "0.0.0.0", "127.255.255.255", 2147483646, "127.255.255.255", 1},
	}
	for _, test := range tests {
		actualNetworkAddress, actualBroadcastAddress, actualNoOfHosts, actualWildCardMask, actualCidrNotation := Calculate(test.IPAddress, test.SubnetMask)
		assert.Equal(t, test.ExpectedNetworkAddress, actualNetworkAddress)
		assert.Equal(t, test.ExpectedBroadcastAddress, actualBroadcastAddress)
		assert.Equal(t, test.ExpectedNoOfHosts, actualNoOfHosts)
		assert.Equal(t, test.ExpectedWildcardMask, actualWildCardMask)
		assert.Equal(t, test.ExpectedCidrNotation, actualCidrNotation)
	}
}

func TestIsValidSubnetMask(t *testing.T) {
	var tests = []struct {
		SubnetMask      string
		ExpectedIsValid bool
	}{
		{"255.255.255.128", true},
		{"255.255.128.0", true},
		{"255.128.255.0", false},
		{"0.0.128.0", false},
	}
	for _, test := range tests {
		maskInt := binary.BigEndian.Uint32(net.ParseIP(test.SubnetMask)[12:16])
		assert.Equal(t, test.ExpectedIsValid, IsValidSubnetMask(maskInt))
	}
}

package subnet

import (
	"encoding/binary"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testAssertionHelper(
	t *testing.T,
	trueValues []interface{},
	networkAddress string,
	broadcastAddress string,
	noOfHosts int,
	wildcardMask string,
	cidrNotation int,
) {
	assert.Equal(t, trueValues[0], networkAddress)
	assert.Equal(t, trueValues[1], broadcastAddress)
	assert.Equal(t, trueValues[2], noOfHosts)
	assert.Equal(t, trueValues[3], wildcardMask)
	assert.Equal(t, trueValues[4], cidrNotation)
}

func TestSubnetCalculateClassC(t *testing.T) {
	// testing subnet for class C
	ip := "198.168.1.101"
	mask := "255.255.255.128"

	trueValues := []interface{}{
		"198.168.1.0", "198.168.1.127", 126, "0.0.0.127", 25,
	}

	networkAddress, broadcastAddress, noOfHosts, wildcardMask, cidrNotation := Calculate(ip, mask)
	testAssertionHelper(t, trueValues, networkAddress, broadcastAddress, noOfHosts, wildcardMask, cidrNotation)
}

func TestSubnetCalculateClassB(t *testing.T) {
	// testing subnet for class B
	ip := "198.168.1.101"
	mask := "255.255.192.0"

	trueValues := []interface{}{
		"198.168.0.0", "198.168.63.255", 16382, "0.0.63.255", 18,
	}

	networkAddress, broadcastAddress, noOfHosts, wildcardMask, cidrNotation := Calculate(ip, mask)
	testAssertionHelper(t, trueValues, networkAddress, broadcastAddress, noOfHosts, wildcardMask, cidrNotation)
}

func TestSubnetCalculateClassA(t *testing.T) {
	// testing subnet for class A
	ip := "198.168.1.101"
	mask := "255.248.0.0"

	trueValues := []interface{}{
		"198.168.0.0", "198.175.255.255", 524286, "0.7.255.255", 13,
	}

	networkAddress, broadcastAddress, noOfHosts, wildcardMask, cidrNotation := Calculate(ip, mask)
	testAssertionHelper(t, trueValues, networkAddress, broadcastAddress, noOfHosts, wildcardMask, cidrNotation)
}

func TestSubnetCalculateTest1(t *testing.T) {
	ip := "52.12.12.52"
	mask := "128.0.0.0"

	trueValues := []interface{}{
		"0.0.0.0", "127.255.255.255", 2147483646, "127.255.255.255", 1,
	}

	networkAddress, broadcastAddress, noOfHosts, wildcardMask, cidrNotation := Calculate(ip, mask)
	testAssertionHelper(t, trueValues, networkAddress, broadcastAddress, noOfHosts, wildcardMask, cidrNotation)
}

func TestSubnetCalculateTest2(t *testing.T) {
	ip := "52.12.12.52"
	mask := "128.0.0.0"

	trueValues := []interface{}{
		"0.0.0.0", "127.255.255.255", 2147483646, "127.255.255.255", 1,
	}

	networkAddress, broadcastAddress, noOfHosts, wildcardMask, cidrNotation := Calculate(ip, mask)
	testAssertionHelper(t, trueValues, networkAddress, broadcastAddress, noOfHosts, wildcardMask, cidrNotation)
}

func TestIsValidSubnetMask1(t *testing.T) {
	subnetMask := "255.255.255.128"
	maskInt := binary.BigEndian.Uint32(net.ParseIP(subnetMask)[12:16])
	assert.Equal(t, true, IsValidSubnetMask(maskInt))
}

func TestIsValidSubnetMask2(t *testing.T) {
	subnetMask := "255.255.128.0"
	maskInt := binary.BigEndian.Uint32(net.ParseIP(subnetMask)[12:16])
	assert.Equal(t, true, IsValidSubnetMask(maskInt))
}

func TestIsValidSubnetMask3(t *testing.T) {
	subnetMask := "255.128.255.0"
	maskInt := binary.BigEndian.Uint32(net.ParseIP(subnetMask)[12:16])
	assert.Equal(t, false, IsValidSubnetMask(maskInt))
}

func TestIsValidSubnetMask4(t *testing.T) {
	subnetMask := "0.0.128.0"
	maskInt := binary.BigEndian.Uint32(net.ParseIP(subnetMask)[12:16])
	assert.Equal(t, false, IsValidSubnetMask(maskInt))
}

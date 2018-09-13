package string_utils

import "fmt"

// IPAddr implements Stringer interface so it can be easily printed.
type IPAddr [4]byte

// String method is provided to implicitly implement the Stringer interface to IPAddr type.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

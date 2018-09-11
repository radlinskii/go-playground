package string_utils

import "fmt"

// IPAddr implements Stringer interface so it can be easily printed.
type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

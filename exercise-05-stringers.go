package tutorial

import "fmt"

// IPAddr is an exported IP address type
type IPAddr [4]byte

// A String() method for IPAddr
func (p IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", p[0], p[1], p[2], p[3])
}

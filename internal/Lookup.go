package internal

import (
	"github.com/muonsoft/validation/validate"
	"net"
)

func PerformLookup(host string) ([]string, error) {
	var names []string
	var err error
	if validate.IP(host) != nil {
		names, err = net.LookupHost(host)
	} else {
		names, err = net.LookupAddr(host)
	}
	return names, err
}

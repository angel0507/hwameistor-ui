package types

import (
	"net"
	"strconv"
)

const (
	DefaultAPIPort = 9500
)

func GetAPIServerAddressFromIP(ip string) string {
	return net.JoinHostPort(ip, strconv.Itoa(DefaultAPIPort))
}

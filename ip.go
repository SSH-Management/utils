package utils

import "net"

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, address := range addrs {
		ipnet, ok := address.(*net.IPNet)

		// check the address type and if it is not a loopback the display it
		if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}

	return ""
}

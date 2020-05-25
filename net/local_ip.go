package net

import (
	"net"
	"os"
)

func GetLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		//uflog.DEBUG(err)
		os.Exit(1)
	}

	for _, address := range addrs {
		//检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

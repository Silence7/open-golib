package net

import (
	"fmt"
	"strconv"
	"strings"
)

func IntToIP(ipnr uint) string {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)
	return fmt.Sprintf("%v.%v.%v.%v", bytes[3], bytes[2], bytes[1], bytes[0])
}

func IPToInt(ip string) (uint, error) {
	var intIP uint

	ips := strings.Split(ip, ".")
	if len(ips) != 4 {
		return 0, fmt.Errorf("ip:%s is not a IP", ip)
	}

	for k, v := range ips {
		i, err := strconv.Atoi(v)
		if err != nil || i > 255 {
			return 0, err
		}
		intIP = intIP | uint(i)<<uint(8*(3-k))
	}

	return intIP, nil
}

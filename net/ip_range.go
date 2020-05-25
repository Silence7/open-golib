package net

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
)

func IpRange2CIDR(startIP, endIP string) ([]string, error) {
	var result = make([]string, 0)
	start, err := IPString2Uint(strings.Trim(startIP, " "))
	if nil != err {
		return nil, fmt.Errorf("invalid start ip")
	}

	end, err := IPString2Uint(strings.Trim(endIP, " "))
	if nil != err {
		return nil, fmt.Errorf("invalid end ip")
	}

	for end >= start {
		maxSize := 32
		for maxSize > 0 {
			mask := iMask(maxSize - 1)
			maskBase := start & mask

			if maskBase != start {
				break
			}
			maxSize = maxSize - 1
		}

		tmp := math.Log(float64(end-start+1)) / math.Log(2)
		maxDiff := int(32 - math.Floor(tmp))
		if maxSize < maxDiff {
			maxSize = maxDiff
		}

		ip, err := Uint2IPString(start)
		if nil != err {
			return nil, fmt.Errorf("convert uint ip to string error:%v", start)
		}

		result = append(result, ip+"/"+strconv.Itoa(maxSize))
		start += uint64(math.Pow(2, float64(32-maxSize)))
	}

	return result, nil
}

func CIDR2IpRange(cidr string) (string, string, error) {
	var beginIp string
	var endIp string

	return beginIp, endIp, nil
}

func iMask(s int) uint64 {
	return uint64(math.Round(math.Pow(2, 32) - math.Pow(2, float64(32-s))))
}

func IPString2Uint(ip string) (uint64, error) {
	b := net.ParseIP(ip).To4()
	if b == nil {
		return 0, fmt.Errorf("invalid ipv4 format")
	}

	return uint64(b[3]) | uint64(b[2])<<8 | uint64(b[1])<<16 | uint64(b[0])<<24, nil
}

func Uint2IPString(i uint64) (string, error) {
	if i > math.MaxUint32 {
		return "", fmt.Errorf("beyond the scope of ipv4")
	}

	ip := make(net.IP, net.IPv4len)
	ip[0] = byte(i >> 24)
	ip[1] = byte(i >> 16)
	ip[2] = byte(i >> 8)
	ip[3] = byte(i)

	return ip.String(), nil
}

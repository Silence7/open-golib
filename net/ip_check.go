package net

import (
	"fmt"
	"net"
	"strings"
)

func CheckIPVaild(ip string) error {
	err := fmt.Errorf("ip:%s invaild", ip)

	if nil == net.ParseIP(ip) {
		return err
	}

	return nil
}

// 检查IP有效,只支持IPv4
func CheckIPv4Vaild(ip string) error {
	if strings.Contains(ip, ":") {
		return fmt.Errorf("ip:%s is not ipv4", ip)
	}

	if nil == net.ParseIP(ip) {
		return fmt.Errorf("ip:%s invaild", ip)
	}

	return nil
}

// 检查IP有效,只支持IPv6
func CheckIPv6Vaild(ip string) error {
	if !strings.Contains(ip, ":") {
		return fmt.Errorf("ip:%s is not ipv6", ip)
	}

	if nil == net.ParseIP(ip) {
		return fmt.Errorf("ip:%s invaild", ip)
	}

	return nil
}

// 检查IP有效,只支持IPv4，可填掩码 1.1.1.1/32
func CheckIPv4MaskVaild(ip string) error {
	if strings.Contains(ip, ":") {
		return fmt.Errorf("ip:%s is not ipv4", ip)
	}

	_, _, err := net.ParseCIDR(ip)

	if nil != err && nil == net.ParseIP(ip) {
		return fmt.Errorf("ip:%s is not ipv4 or ipv4/mask", ip)
	}

	return nil
}

// 检查端口是否有效
func CheckPortIsVaild(port int) error {
	if (port < 0) || (port > 65535) {
		return fmt.Errorf("port:%d invaild", port)
	}

	return nil
}

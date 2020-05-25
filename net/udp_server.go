package net

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"net"
	"time"
)

//PacketHandle ...
type PacketHandle func([]byte, int) (bool, error)

//UDPServerBase ...
type UDPServerBase struct {
	Name       string
	IPStr      string
	Port       int
	Sock       *net.UDPConn
	BufferSize int
	HandleList []PacketHandle //报文处理列表
	LegalPeer  []string       //可信的对端
	TimeOut    time.Duration  //超时时间
}

//UDPServerStats ...
type UDPServerStats struct {
	Rx NetEventStat
	Tx NetEventStat
}

//UDPServer ...
type UDPServer struct {
	Base  UDPServerBase
	Stats UDPServerStats
}

//AddPktHandle ...
func (ptr *UDPServer) AddPktHandle(handle PacketHandle) {
	ptr.Base.HandleList = append(ptr.Base.HandleList, handle)
}

func (ptr *UDPServer) String() string {
	return fmt.Sprintf("UDP Server Name:%s IP:%s Port:%d", ptr.Base.Name, ptr.Base.IPStr, ptr.Base.Port)
}

//Process ...
func (ptr *UDPServer) Process(data []byte, n int) {
	for idx, handle := range ptr.Base.HandleList {
		result, err := handle(data, n)
		if err != nil {
			logs.Error("%s handle %d execute error %s", ptr.String(), idx, err.Error())
			continue
		}
		if !result {
			logs.Error("%s handle %d execute faield", ptr.String(), idx)
			continue
		}
	}
}

//NewUDPServer ...
func NewUDPServer(name string, ip string, port int, timeout time.Duration) *UDPServer {
	out := new(UDPServer)
	if nil == out {
		return nil
	}
	out.Base.Name = name
	out.Base.IPStr = ip
	out.Base.Port = port
	out.Base.TimeOut = timeout
	out.Base.BufferSize = 2048
	sock, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP(out.Base.IPStr), Port: out.Base.Port})
	if err != nil {
		logs.Error("init udp server ip:%s port:%d failed err :%s", out.Base.IPStr, out.Base.Port, err.Error())
		return nil
	}
	if nil == sock {
		logs.Error("init udp server ip:%s port:%d failed get nil sock", out.Base.IPStr, out.Base.Port)
		return nil
	}
	out.Base.Sock = sock
	return out
}

//RecvTask ...
func (ptr *UDPServer) RecvTask() {
	logs.Info("RecvTask for %s", ptr.String())
	buffer := make([]byte, ptr.Base.BufferSize)

	for true {
		n, remoteAddr, err := ptr.Base.Sock.ReadFromUDP(buffer)
		ptr.Stats.Rx.Total.AddStats(0, 1)
		if err != nil {
			logs.Debug("%s during read: %s , %s", ptr.String(), err, remoteAddr)
			ptr.Stats.Rx.Error.AddStats(0, 1)
			continue
		}
		ptr.Stats.Rx.Valid.AddStats(uint64(n), 1)
		ptr.Process(buffer[:n:n], n)
	}
}

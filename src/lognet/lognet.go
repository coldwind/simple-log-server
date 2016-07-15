package lognet

import (
	"fmt"
	"net"
	"os"

	"github.com/c4pt0r/ini"
)

type NetModel struct {
	UdpConn *net.UDPConn
}

func (this *NetModel) Run() {

	conf := ini.NewConf("base.ini")
	serverIp := conf.String("sys", "ip", "")
	serverPort := conf.String("sys", "port", "")
	conf.Parse()

	serverAddr := *serverIp + ":" + *serverPort

	udpAddr, err := net.ResolveUDPAddr("udp4", serverAddr)
	this.checkErr(err)

	this.UdpConn, err = net.ListenUDP("udp4", udpAddr)
	this.checkErr(err)
}

func (this *NetModel) checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

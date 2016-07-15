// xtrops_log_server project main.go
package main

import (
	"fmt"
	"logfile"
	"lognet"
)

func main() {
	lognetObj := &lognet.NetModel{}
	lognetObj.Run()
	defer lognetObj.UdpConn.Close()

	logfileObj := &logfile.FileModel{}
	logfileObj.Init()

	readBuf := make([]byte, 1024)
	for {
		_, _, err := lognetObj.UdpConn.ReadFromUDP(readBuf)
		if err != nil {
			continue
		}
		fmt.Println(string(readBuf))
		go logfileObj.Record(readBuf)
	}
}

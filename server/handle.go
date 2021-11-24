package server

import (
	"net"
	"resonance/logger"
	"resonance/util"
)

var ConnMap = make(map[string]net.Conn)

func HandleConn(conn net.Conn) {
	addr := conn.RemoteAddr().String()

	ConnMap[addr] = conn
	defer func() {
		conn.Close()
		delete(ConnMap, addr)
	}()

	logger.Info(addr + " conncet sucessful")

	util.SendFile(*MusicName, conn)

	select {}
}

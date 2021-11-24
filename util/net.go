package util

import (
	"net"
	"resonance/logger"
)

func WriteToConn(message string, conn net.Conn) {
	_, err := conn.Write([]byte(message))
	if err != nil {
		logger.Error("The message is \"%s\"", message)
		logger.Error("Failed to write message to connect" + err.Error())
		return
	}
}

func ReadFromConn(conn net.Conn) (string, int, []byte) {
	buf := make([]byte, 1024*4)
	n, err := conn.Read(buf)
	if err != nil {
		logger.Error("Failed to read message from connect " + err.Error())
	}
	return string(buf[:n]), n, buf
}

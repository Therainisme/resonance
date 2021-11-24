package client

import (
	"flag"
	"net"
	"resonance/logger"
	"resonance/util"
)

// Here is the address of my campus network
// I'm too lazy to alter
var serverAddress = flag.String("serverAddress", "10.33.60.104:8000", "server address")

func Run() {
	// Connect to server
	conn, err := net.Dial("tcp", *serverAddress)
	if err != nil {
		logger.Error("Dial err:", err)
		return
	}
	defer conn.Close()

	playName := util.ReceiveFile(conn)

	for {
		message, _, _ := util.ReadFromConn(conn)

		logger.Info("%s", message)
		if message == "PLAY" {
			util.Play(playName)
		}
	}
}

package server

import (
	"flag"
	"net"
	"os"
	"resonance/logger"
	"resonance/util"
)

var MusicName = flag.String("musicName", "", "music name")

func Run() {
	// Listening
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		logger.Error("err = ", err)
		return
	}
	defer listener.Close()

	// Create go coroutine to receive keyboard input
	go func() {
		buf := make([]byte, 1024*4)
		for {
			n, _ := os.Stdin.Read(buf)
			message := string(buf[:n-2])

			if message == "PLAY" {
				logger.Info("START PLAYING")
				go util.Play(*MusicName)
			}

			for _, conn := range ConnMap {
				util.WriteToConn(message, conn)
			}
		}
	}()

	// Handle multiple connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("err = ", err)
			return
		}

		go HandleConn(conn)
	}
}

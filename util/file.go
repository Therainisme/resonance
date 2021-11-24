package util

import (
	"io"
	"net"
	"os"
	"resonance/logger"
	"time"
)

func SendFile(fileName string, conn net.Conn) {
	// 1. Send the file name to the receiver
	WriteToConn(fileName, conn)

	// 2. If client reply "FILE_NAME_ACK", the recipient is ready to receive the file
	response, _, _ := ReadFromConn(conn)
	if response != "FILE_NAME_ACK" {
		logger.Error("Incorrect confirmation message received")
		return
	}

	// 3. After receiving the confirmation, start sending file
	logger.Info("Start sending files to " + conn.RemoteAddr().String())

	// 4. Open the named file for reading.
	f, err := os.Open(fileName)
	if err != nil {
		logger.Error("Failed to open file" + err.Error())
		return
	}
	defer f.Close()

	// 5. Read 4096 each time until the end of the file
	buf := make([]byte, 1024*4)
	for {
		n, err := f.Read(buf)

		if err == io.EOF {
			// 6. If EOF, send an end message
			logger.Info("EOF, Wait for 3 seconds")
			time.Sleep(time.Second * 3)

			// 7. Notify the client that sending a file is complete
			WriteToConn("FILE_END\r\n", conn)
			return
		}

		// 5.5 Send once
		conn.Write(buf[:n])
	}
}

func ReceiveFile(conn net.Conn) string {
	buf := make([]byte, 1024*4)

	// 1. 读取对方发送的文件名
	fileName, _, _ := ReadFromConn(conn)

	// 2. 回复"FILE_NAME_ACK"
	WriteToConn("FILE_NAME_ACK", conn)

	//新建文件
	f, _ := os.Create(fileName)

	// 3. 接受对方发送过来的文件内容
	logger.Info("Start to receive " + fileName)
	for {
		n, err4 := conn.Read(buf)
		if err4 != nil {
			logger.Error("conn.Read err" + err4.Error())
			return ""
		}

		if string(buf[:n]) == "FILE_END" || string(buf[:n-1]) == "FILE_END" || string(buf[:n-2]) == "FILE_END" {
			logger.Info("Complete")
			return fileName
		}
		f.Write(buf[:n])
	}
}

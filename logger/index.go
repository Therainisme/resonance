package logger

import (
	"fmt"
	"time"
)

const (
	colorRed = uint8(iota + 91)
	colorGreen
	colorYellow
	colorBlue
	colorMagenta

	info = "[INFO]"
	trac = "[TRAC]"
	erro = "[ERRO]"
	warn = "[WARN]"
	succ = "[SUCC]"
)

func Error(format string, a ...interface{}) {
	prefix := red(erro)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a...))
}

func Info(format string, a ...interface{}) {
	prefix := blue(info)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a...))
}

func formatLog(prefix string) string {
	return time.Now().Format("2006-01-02 15:04:05") + " " + prefix
}

func blue(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorBlue, s)
}

func red(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorRed, s)
}

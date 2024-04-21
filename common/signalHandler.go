package common

import (
	"os"
	"os/signal"
	"syscall"
)

var exit chan os.Signal

func SetupSignalHandler() chan os.Signal {
	if exit != nil {
		return exit
	}

	exit = make(chan os.Signal, 4)
	signal.Notify(exit, os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	return exit
}

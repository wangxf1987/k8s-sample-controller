package signals

import (
	"os"
	"os/signal"
)

var onlyOneSignalHandler = make(chan struct{})

func SetupSignalHandler() (stopCh <-chan struct{}) {
	close(onlyOneSignalHandler) // panic when called twice

	stop := make(chan struct{})
	c := make(chan os.Signal)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1)
	}()

	return stop
}

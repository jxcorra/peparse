package status

import (
	"os"
	"os/signal"
	"syscall"
)

func WatchDone(communication chan bool) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	func() {
		<-sigs
		communication <- true
	}()
}

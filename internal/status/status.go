package status

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func WatchDone(communication chan bool, wg *sync.WaitGroup) {
	wg.Add(1)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	func() {
		defer wg.Done()
		<-sigs
		communication <- true
	}()
}

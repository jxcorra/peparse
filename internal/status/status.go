package status

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func WatchTermination(communication chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	communication <- true
}

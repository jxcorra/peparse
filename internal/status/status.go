package status

import (
	"sync"

	"github.com/jxcorra/peparse/internal/common"
)

func WatchTermination(communication *common.Communication, wg *sync.WaitGroup) {
	defer wg.Done()
	<-communication.Signals
	communication.Done <- true
}

package status_test

import (
	"os/signal"
	"sync"
	"syscall"
	"testing"

	"github.com/jxcorra/peparse/internal/config"
	"github.com/jxcorra/peparse/internal/status"
)

func TestWatchTermination(t *testing.T) {
	type testCase struct {
		signalStr string
		signal    syscall.Signal
	}

	testCases := []testCase{
		{
			signalStr: "sigint",
			signal:    syscall.SIGINT,
		},
		{
			signalStr: "sigterm",
			signal:    syscall.SIGTERM,
		},
	}

	for _, testCase := range testCases {
		communication := config.NewCommunication(1)
		signal.Notify(communication.Signals, syscall.SIGINT, syscall.SIGTERM)
		var wg sync.WaitGroup

		wg.Add(1)
		go status.WatchTermination(communication, &wg)

		wg.Add(1)
		go func() {
			defer wg.Done()
			err := syscall.Kill(syscall.Getpid(), testCase.signal)
			if err != nil {
				t.Errorf("error killing process %d", syscall.Getpid())
			}
		}()

		wg.Wait()

		if len(communication.Done) != 1 {
			t.Errorf("not done with signal `%s`", testCase.signalStr)
		}
	}
}

package status_test

import (
	"sync"
	"syscall"
	"testing"

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
		done := make(chan bool, 1)
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			err := syscall.Kill(syscall.Getpid(), testCase.signal)
			if err != nil {
				t.Errorf("error killing process %d", syscall.Getpid())
			}
		}()

		status.WatchTermination(done, &wg)
		wg.Wait()

		if len(done) != 1 {
			t.Errorf("not done with signal `%s`", testCase.signalStr)
		}
	}
}

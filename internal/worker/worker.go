package worker

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/html"

	common "github.com/jxcorra/peparse/internal/common"
	parse "github.com/jxcorra/peparse/internal/parse"
)

func RunPeriodicTask(parameters common.Parameters) {
	for i := 0; i < parameters.NumOfWorkers; i++ {
		parameters.Wg.Add(1)
		go worker(parameters.Tasks, parameters.Output, parameters.Wg, &parameters.Communication)
	}

	parameters.Wg.Add(1)
	go printer(parameters.Output, parameters.Wg, &parameters.Communication)

	ticker := time.NewTicker(time.Millisecond * time.Duration(parameters.Period))

	for {
		select {
		case <-ticker.C:
			for _, resource := range parameters.Resources.Resources {
				parameters.Tasks <- resource
			}
		case <-parameters.Communication.Done:
			for i := 0; i < parameters.NumOfWorkers; i++ {
				parameters.Communication.WorkerDone <- true
			}
			parameters.Communication.PrinterDone <- true
			return
		}

	}
}

func worker(resources chan common.ResourceConfig, output chan common.Parsed, wg *sync.WaitGroup, communication *common.Communication) {
	defer wg.Done()
	for {
		select {
		case resource, more := <-resources:
			if !more {
				return
			}
			response, err := http.Get(resource.Url)
			if err != nil {
				fmt.Printf("cannot get %s", resource.Url)
				continue
			}

			tokenizer := html.NewTokenizer(response.Body)
			data, err := parse.ParsePage(tokenizer, &resource.Search)
			response.Body.Close()

			if err != nil {
				fmt.Printf("%s %s", resource.Url, err.Error())
				continue
			}

			output <- data
		case workerDone := <-communication.WorkerDone:
			if workerDone {
				return
			}
		}
	}
}

func printer(output chan common.Parsed, wg *sync.WaitGroup, communication *common.Communication) {
	defer wg.Done()
	for {
		select {
		case message, more := <-output:
			if !more {
				return
			}
			serialized, err := parse.Serialize(&message)
			if err != nil {
				fmt.Println(fmt.Errorf(err.Error()))
			}
			fmt.Println(serialized)
		case printerDone := <-communication.PrinterDone:
			if printerDone {
				return
			}
		}

	}
}

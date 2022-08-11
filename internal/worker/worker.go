package worker

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/html"

	common "github.com/jxcorra/peparse/internal/common"
	parse "github.com/jxcorra/peparse/internal/parse"
)

func RunPeriodicTask(parameters common.Parameters) {
	for i := 0; i < parameters.NumOfWorkers; i++ {
		go worker(parameters.Tasks, parameters.Output)
	}

	go printer(parameters.Output)

	for range time.Tick(time.Millisecond * time.Duration(parameters.Period)) {
		for _, resource := range parameters.Resources.Resources {
			parameters.Tasks <- resource
		}
	}
}

func worker(communication chan common.ResourceConfig, output chan common.Parsed) {
	for resource := range communication {
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
	}
}

func printer(output chan common.Parsed) {
	for message := range output {
		serialized, err := parse.Serialize(&message)
		if err != nil {
			fmt.Println(fmt.Errorf(err.Error()))
		}
		fmt.Println(serialized)
	}
}

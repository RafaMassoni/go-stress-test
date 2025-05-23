package runner

import (
	"fmt"
	"net/http"
	"stress-test/config"
	"sync"
	"time"
)

type TestResults struct {
	TotalRequests    int
	Successful       int
	Duration         time.Duration
	StatusCodesCount map[int]int
}

func RunLoadTest(cfg *config.ConfigParams) TestResults {
	fmt.Printf("Executando teste de carga para %s com %d requests e %d chamadas simultâneas...\n",
		cfg.Url, cfg.Requests, cfg.Concurrency)

	var wg sync.WaitGroup
	requestsPerWorker := cfg.Requests / cfg.Concurrency
	extraRequests := cfg.Requests % cfg.Concurrency

	statusCodes := make(chan int, cfg.Requests)
	startTime := time.Now()

	for i := 0; i < cfg.Concurrency; i++ {
		wg.Add(1)

		numRequests := requestsPerWorker
		if i < extraRequests {
			numRequests++
		}

		go func(worker int, num int) {
			defer wg.Done()
			for j := 0; j < num; j++ {
				resp, err := http.Get(cfg.Url)
				if err != nil {
					fmt.Printf("Request %d/%d para %s => 0\n", j, worker, cfg.Url)
					statusCodes <- 0
					continue
				}
				statusCodes <- resp.StatusCode
				fmt.Printf("Request %d/%d para %s => %d\n", j, worker, cfg.Url, resp.StatusCode)
				resp.Body.Close()
			}
		}(i, numRequests)
	}

	wg.Wait()
	close(statusCodes)

	duration := time.Since(startTime)

	resultMap := make(map[int]int)
	successCount := 0

	for code := range statusCodes {
		resultMap[code]++
		if code == 200 {
			successCount++
		}
	}

	return TestResults{
		TotalRequests:    cfg.Requests,
		Duration:         duration,
		Successful:       successCount,
		StatusCodesCount: resultMap,
	}
}

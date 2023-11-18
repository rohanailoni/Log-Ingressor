package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-faker/faker/v4"
)

const numRequests = 100 // Total number of requests to send
const poolSize = 10     // Size of the worker pool

var wg sync.WaitGroup

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	wg.Add(numRequests)
	// Create a channel to coordinate worker pool
	requests := make(chan int, numRequests)

	// Create a worker pool
	for i := 0; i < poolSize; i++ {
		go worker(requests)
	}

	// Enqueue requests
	for i := 0; i < numRequests; i++ {
		requests <- i
	}
	close(requests) // Close the channel to signal workers that no more requests will be sent

	// Wait for all workers to finish
	wg.Wait()
}

func worker(requests chan int) {
	for {
		select {
		case i, ok := <-requests:
			if !ok {
				// Channel closed, no more requests
				return
			}
			sendRequest(i)
		}
	}
}

type ErrorLog struct {
	Level            string `faker:"oneof:error,debug"`
	Message          string `faker:"sentence"`
	ResourceID       string `faker:"uuid_hyphenated"`
	Timestamp        string `faker:"timestamp"`
	TraceID          string `faker:"uuid_hyphenated"`
	SpanID           string `faker:"oneof:spanid,span"`
	Commit           string `faker:"jwt"`
	ParentResourceID string `faker:"uuid_hyphenated"`
}

func generateRandomData() (ErrorLog, error) {
	var errorLog ErrorLog
	err := faker.FakeData(&errorLog)
	if err != nil {
		fmt.Println("Error generating random data:", err)
		return ErrorLog{}, err
	}
	return errorLog, nil
}

func sendRequest(requestNum int) {
	defer wg.Done()

	url := "http://localhost:3000"
	data, err := generateRandomData()
	if err != nil {
		fmt.Println(err)
		return
	}
	data.SpanID = "spanId-" + fmt.Sprintf("%d", rand.Intn(100))
	jsonStr, err := structToJSON(data)
	if err != nil {
		fmt.Println("error while encoding", err)
		return
	}

	method := "POST"
	payload := strings.NewReader(jsonStr)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println("error while creating a new request", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("the error occurred here", err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Request %d: %s\n", requestNum, string(body))
}

func structToJSON(data interface{}) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

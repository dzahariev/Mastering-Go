package main

import (
	"fmt"
	"net/http"
	"time"
)

var timeout = make(chan bool)

func main() {
	defer close(timeout)

	go startServer()

	for {
		// Do the work
		doTheJob()
		// Read the channel and decide to continue or not
		select {
		case <-timeout:
			break
		case <-time.After(time.Second * 10):
			break
		}
	}
}

func doTheJob() {
	fmt.Printf("Hello: %s \n", time.Now().Format(time.Stamp))
	time.Sleep(time.Second * 1)
}

func startServer() {
	http.HandleFunc("/trigger", handleTrigger)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Printf("failed to start the server with error: %v \n", err)
	}
}

func handleTrigger(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	select {
	case timeout <- true:
		w.Write([]byte(`{ "status": "OK" }`))
	default:
		w.Write([]byte(`{ "status": "IGNORED" }`))
	}
}

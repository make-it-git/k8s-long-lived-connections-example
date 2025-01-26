package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var SERVER_ADDR = os.Getenv("SERVER_ADDR")

func main() {
	client := &http.Client{
		Timeout: time.Second,
		// Uncomment to disable keep-alives
		//Transport: &http.Transport{
		//	DisableKeepAlives: true,
		//},
	}

	i := 0
	for {
		i++
		err := sendRequest(client, i)
		if err != nil {
			log.Printf("Error sending request: %v, retrying...", err)
			time.Sleep(time.Millisecond * 100)
			continue
		}
		time.Sleep(time.Second)
	}
}

func sendRequest(client *http.Client, i int) error {
	req, err := http.NewRequest(http.MethodGet, SERVER_ADDR, nil)
	if i%5 == 0 {
		fmt.Println("x-user 1")
		req.Header.Set("x-user", "1")
	} else {
		fmt.Println("x-user 2")
		req.Header.Set("x-user", "2")
	}
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error during request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned non-200 status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading body: %w", err)
	}

	fmt.Printf("Body: %s\n", string(body))
	return nil
}

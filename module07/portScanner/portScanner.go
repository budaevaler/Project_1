package portScanner

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"
)

func worker(ctx context.Context, ports <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case port, ok := <-ports:
			if !ok {
				return
			}
			address := fmt.Sprintf("localhost:%d", port)
			conn, err := net.Dial("tcp", address)
			if err == nil {
				conn.Close()
				results <- port
			}
		}
	}
}

func Scan(startPort int, endPort int, threadCount int, timeoutSec int) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSec)*time.Second)
	defer cancel()

	ports := make(chan int, endPort-startPort+1)
	results := make(chan int, endPort-startPort+1)

	var wg sync.WaitGroup
	for i := 0; i < threadCount; i++ {
		wg.Add(1)
		go worker(ctx, ports, results, &wg)
	}

	go func() {
		defer close(ports)
		for p := startPort; p <= endPort; p++ {
			select {
			case <-ctx.Done():
				return
			case ports <- p:
			}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("Открытые порты:")
	for r := range results {
		fmt.Println(r)
	}
}

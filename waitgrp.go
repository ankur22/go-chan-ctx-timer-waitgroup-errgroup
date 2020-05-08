package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Printf("Caught %s\n", sig)
		cancel()
	}()

	filename := "test.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Can't load %s - %v\n", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make(chan string, 1)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			close(lines)
		}()
		for {
			if !scanner.Scan() {
				fmt.Println("Reader: Completed")
				return
			}
			lines <- scanner.Text()

			select {
			case <-ctx.Done():
				fmt.Println("Reader: Context closed")
				return
			default:
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("reader error - %v\n", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Sender: Context closed")
				return
			case l, ok := <-lines:
				if !ok {
					fmt.Printf("Sender: Channel closed\n", l)
					return
				}
				fmt.Printf("Sender: Sending %s to remote database\n", l)
			}
		}
	}()

	wg.Wait()
}

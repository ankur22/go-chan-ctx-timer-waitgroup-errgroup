package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
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

	go func() {
		fmt.Println("Reading from file")
		<-ctx.Done()
		fmt.Println("Reader: Context closed")
	}()

	go func() {
		fmt.Println("Sending to remote database")
		<-ctx.Done()
		fmt.Println("Sender: Context closed")
	}()

	<-ctx.Done()
	fmt.Println("Main: Context closed")
}

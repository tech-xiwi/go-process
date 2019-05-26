package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)

	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		_ = http.ListenAndServe(":1234", nil)
	}()
	go func() {
		sig := <-sigs
		fmt.Println("algo sig:", sig)
		done <- true
	}()
	fmt.Println("algo wait signal")
	<-done
	fmt.Println("algo exit")
}

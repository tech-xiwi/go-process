package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)
//golang 进程热重启
//https://cloud.tencent.com/developer/article/1388556
func main() {
	sigs := make(chan os.Signal, 1)

	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		_ = http.ListenAndServe(":8100", nil)
	}()
	go func() {
		cmd := "D:\\workspace\\golang\\src\\studio.xiwi\\process\\algo\\algo.exe"
		p, err := os.StartProcess(cmd, []string{cmd}, &os.ProcAttr{
			Dir:"D:\\workspace\\golang\\src\\studio.xiwi\\process\\algo",
			Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		})
		if err != nil {
			panic(err)
		}
		ps ,err := p.Wait()
		if err != nil {
			panic(err)
		}
		fmt.Println(ps.Pid())
	}()
	go func() {
		sig := <-sigs
		fmt.Println("sig:", sig)
		done <- true
	}()
	fmt.Println("wait signal")
	<-done
	fmt.Println("exit")
}

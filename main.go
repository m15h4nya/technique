package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/m15h4nya/online_vault/common"
)

func main() {
	wg := sync.WaitGroup{}
	server := NewServer()

	exit := common.SetupSignalHandler()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-exit
		fmt.Println("Shutting down")

		err := server.Shutdown(context.Background())
		if err != nil {
			fmt.Println("Shutdown error: ", err)
		}
	}()

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("ListenAndServer error: ", err)

	}

	wg.Wait()
}

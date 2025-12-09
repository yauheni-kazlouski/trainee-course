package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Cake struct {
	BakedBy  int
	BakeTime int
	PackedBy int
	PackTime int
}

func main() {
	
	n := 5
	m := 15
	k := 1000
	t1 := 2
	t2 := 1

	bakerChan := make(chan struct{}, k)
	packerChan := make(chan Cake, k)
	result := make(chan Cake, k)
	done := make(chan struct{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(){
		defer close(bakerChan)
		for range k{
			select{
			case <-ctx.Done():
				return
			default:
				bakerChan <- struct{}{}
			}
		}
	}()
	
	bakerWg := sync.WaitGroup{}
	bakerWg.Add(n)
	for i := range(n) {
		go func (workerId int) {
			defer bakerWg.Done()
			for {
				select{
				case _, ok := <- bakerChan:
					if ok{
						time.Sleep(time.Duration(workerId + t1) * time.Millisecond)
						packerChan <- Cake{BakedBy: workerId, BakeTime: workerId + t1}
					} else {
						return
					}
				case <- ctx.Done():
					return
				}
			}
		}(i)
	}
	go func() {
		bakerWg.Wait()
		close(packerChan)
	}()

	packerWg := sync.WaitGroup{}
	packerWg.Add(m)
	for i := range(m) {
		go func (workerId int) {
			defer packerWg.Done()
			for {
				select{
				case cake, ok := <- packerChan:
					if !ok {
						return
					}
					time.Sleep(time.Duration(workerId + t2) * time.Millisecond)
					cake.PackedBy = workerId
					cake.PackTime = workerId + t2
					result <- cake
				case <-ctx.Done():
					return
				}
			}
		}(i)
	}
	go func() {
		packerWg.Wait()
		close(result)
	}()

	// I think there's no need to select ctx.Done in this gorouting, 
	// because shutting down gracefully programm must print all the cakes 
	// that have been already done and are waiting in the result chan buffer to be displayed
	go func() {
		defer close(done)
		for cake := range result{
			fmt.Printf("%#v\n", cake)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	select {
	case <-done:
		fmt.Printf("\nprogramm finished successfully\n")
	case <-sigChan:
		fmt.Printf("\nprogramm interrupted, shutting down...\n\n")
		cancel()
		<-done
	}
	
	fmt.Println("\nprogram ended")
}
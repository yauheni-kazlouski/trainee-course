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

type BakersPool struct{
	jobs	chan struct{}
	result 	chan Cake
	t1		int
	wg		sync.WaitGroup
}

func NewBakersPool(numBakers, numJobs int) *BakersPool {
	pool := &BakersPool{
		jobs: make(chan struct{}, numJobs),
		result: make(chan Cake, numBakers),
		t1: 2,
	}

	pool.wg.Add(numBakers)
	for i := range(numBakers){
		go pool.worker(i)
	}
	return pool
}

func (pool *BakersPool) worker(i int) {
	defer pool.wg.Done()
	for range pool.jobs {
		time.Sleep(time.Duration(i + pool.t1) * time.Millisecond)
		pool.result <- Cake{BakedBy: i, BakeTime: i + pool.t1}
	}
}

func (pool *BakersPool) AddJob(){
	pool.jobs <- struct{}{}
}

func (pool *BakersPool) Close() {
	close(pool.jobs)
	pool.wg.Wait()
	close(pool.result)
}





type PackersPool struct{
	jobs	chan Cake
	result 	chan Cake
	t2		int
	wg		sync.WaitGroup
}

func NewPackersPool(numPackers, numJobs int) *PackersPool {
	pool := &PackersPool{
		jobs: make(chan Cake, numJobs),
		result: make(chan Cake, numPackers), 
		t2: 3,
	}

	pool.wg.Add(numPackers)
	for i := range numPackers {
		go pool.worker(i)
	}

	return pool
}

func (pool *PackersPool) worker(i int) {
	defer pool.wg.Done()

	for cake := range pool.jobs {
		time.Sleep(time.Duration(i + pool.t2) * time.Millisecond)
		cake.PackedBy = i
		cake.PackTime = i + pool.t2
		pool.result <- cake
	}
}

func (pool *PackersPool) AddJob(cake Cake){
	pool.jobs <- cake
}

func (pool *PackersPool) Close() {
	close(pool.jobs)
	pool.wg.Wait()
	close(pool.result)
}

func main() {
	
	n := 5
	m := 10
	k := 10000
	var poolsWG sync.WaitGroup
	
	bakers := NewBakersPool(n, k)
	packers := NewPackersPool(m, k)

	ctx, cancel := context.WithCancel(context.Background())
	
	poolsWG.Add(1)
	go func(){
		defer poolsWG.Done()
		defer bakers.Close()
		for range k {
			select{
			case <-ctx.Done():
				return
			default:
				bakers.AddJob()
			}
		}
	}()

	poolsWG.Add(1)
	go func(){
		defer poolsWG.Done()
		defer packers.Close()

		for cake := range bakers.result {
			select{
			case <-ctx.Done():
				return
			default:
				packers.AddJob(cake)
			}
		}
	}()

	poolsWG.Add(1)
	go func(){
		for cake := range packers.result {
			fmt.Printf("%#v\n", cake)
		}
		poolsWG.Done()
	}()

	sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan struct{})

	go func(){
		poolsWG.Wait()
		close(done)
	}()

	select {
	case <-sigChan:
		fmt.Println("shutting down, waiting to finish goroutings")
	case <-done:
		fmt.Println("All done")
	}

	cancel()
	poolsWG.Wait()
	fmt.Println("programm is finished")
}
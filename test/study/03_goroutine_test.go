package study

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// 协程同步机制
func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		_, _ = fmt.Println("123123")
	}(&wg)
	go func(wg *sync.WaitGroup) {
		wg.Add(2)
		_, _ = fmt.Println("qwe")
	}(&wg)
	go func() {
		wg.Done()
		_, _ = fmt.Println("done ...")
	}()
	go func() {
		wg.Done()
		_, _ = fmt.Println("done ...")
	}()
	go func() {
		wg.Done()
		_, _ = fmt.Println("done ...")
	}()
	wg.Wait()
	_, _ = fmt.Println("all done ...")
}

// 开启协程
func TestGoroutine(t *testing.T) {
	_, _ = fmt.Println(runtime.NumCPU()) // CPU 核心数

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	go send(ch1, ch2, ch3)
	go receive(ch1, ch2, ch3)

	time.Sleep(time.Millisecond)
}

// 并发计算范围内的所有质数
func TestPrimes(t *testing.T) {
	goal := 100
	_, _ = fmt.Println("goal=", goal)
	ch := make(chan int)
	defer close(ch)
	go primeTask(ch)
	for i := 2; i <= goal; i++ {
		ch <- i
	}
}

func primeTask(ch chan int) {
	p := <-ch
	_, _ = fmt.Println(p)
	tmpCh := make(chan int)
	go primeTask(tmpCh)
	for {
		i := <-ch
		if i%p != 0 {
			tmpCh <- i
		}
	}
}

func send(ch1, ch2, ch3 chan<- string) {
	for i := 0; i < 10; i++ {
		ch1 <- fmt.Sprintf("channel 1 : %d", i)
		ch2 <- fmt.Sprintf("channel 2 : %d", i)
		ch3 <- fmt.Sprintf("channel 3 : %d", i)
	}
}

func receive(ch1, ch2, ch3 <-chan string) {
	for {
		select {
		case v := <-ch1:
			_, _ = fmt.Println(v)
		case v := <-ch2:
			_, _ = fmt.Println(v)
		case v := <-ch3:
			_, _ = fmt.Println(v)
		}
	}
}

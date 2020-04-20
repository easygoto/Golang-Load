package test

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"

	. "MergeSort"
)

// 排序演示
func TestArraySource(t *testing.T) {
	result := Merge(
		InMemorySort(
			ArraySource(2234, 8536, 23, 8746, 3486, 412, 3423, 4536)),
		InMemorySort(
			ArraySource(2345, 456, 483, 929, 0404, 367, 987, 8796)))
	for v := range result {
		_, _ = fmt.Printf("%d ", v)
	}
	_, _ = fmt.Println()
}

// 写随机数据到文件
func TestWriteData(t *testing.T) {
	const num = 100
	filename := "../small.dat"
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := RandomSource(num)
	writer := bufio.NewWriter(file)
	WriterSink(writer, p)
	_ = writer.Flush()
}

// 从文件读数据
func TestReadData(t *testing.T) {
	filename := "../small.sort.dat"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := ReaderSource(bufio.NewReader(file), -1)
	for v := range p {
		_, _ = fmt.Println(v)
	}
}

// 单机并发排序
func TestCreatePipeline(t *testing.T) {
	filename, fileSize, chunkCount := "../small.dat", 800, 8
	chunkSize := fileSize / chunkCount
	Init()

	var sortResult []<-chan int
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		_, _ = file.Seek(int64(i*chunkSize), 0)
		source := ReaderSource(bufio.NewReader(file), chunkSize)
		sortResult = append(sortResult, InMemorySort(source))
	}

	p := MergeN(sortResult...)
	writeToFile(p, "../small.sort.dat")
}

// 基于 TCP 传输多机并发排序
func TestCreateNetworkPipeline(t *testing.T) {
	filename, fileSize, chunkCount := "../small.dat", 800, 8
	chunkSize := fileSize / chunkCount
	Init()

	var sortAddr []string
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)

		source := ReaderSource(bufio.NewReader(file), chunkSize)

		addr := ":" + strconv.Itoa(7000+i)
		NetworkSink(addr, InMemorySort(source))
		sortAddr = append(sortAddr, addr)
	}

	var sortResult []<-chan int
	for _, addr := range sortAddr {
		sortResult = append(sortResult, NetworkSource(addr))
	}
	p := MergeN(sortResult...)
	writeToFile(p, "../small.sort.dat")
}

// 通道写数据到文件
func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	WriterSink(writer, p)
}

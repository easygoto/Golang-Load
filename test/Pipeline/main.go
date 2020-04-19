package main

import (
    "Pipeline"
    "bufio"
    "fmt"
    "os"
    "strconv"
)

var filename = "small.dat"

func main() {
    //p := CreatePipeline(filename, 800000000, 8)
    //writeToFile(p, "source.sort.dat")
    //printFile("source.sort.dat")

    //WriteData()

    p := CreateNetworkPipeline("small.dat", 8000, 8)
    writeToFile(p, "small.sort.dat")
    printFile("small.sort.dat")
}

func printFile(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    p := Pipeline.ReaderSource(bufio.NewReader(file), -1)
    count := 0
    for v := range p {
        count++
        _, _ = fmt.Println(v)
        if count >= 0 {
            break
        }
    }
}

func writeToFile(p <-chan int, filename string) {
    file, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    writer := bufio.NewWriter(file)
    defer writer.Flush()

    Pipeline.WriterSink(writer, p)
}

func CreatePipeline(filename string, fileSize, chunkCount int) <-chan int {
    chunkSize := fileSize / chunkCount
    Pipeline.Init()

    var sortResult []<-chan int
    for i := 0; i < chunkCount; i++ {
        file, err := os.Open(filename)
        if err != nil {
            panic(err)
        }

        file.Seek(int64(i*chunkSize), 0)

        source := Pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
        sortResult = append(sortResult, Pipeline.InMemorySort(source))
    }

    return Pipeline.MergeN(sortResult...)
}

func CreateNetworkPipeline(filename string, fileSize, chunkCount int) <-chan int {
    chunkSize := fileSize / chunkCount
    Pipeline.Init()

    var sortAddr []string
    for i := 0; i < chunkCount; i++ {
        file, err := os.Open(filename)
        if err != nil {
            panic(err)
        }

        file.Seek(int64(i*chunkSize), 0)

        source := Pipeline.ReaderSource(bufio.NewReader(file), chunkSize)

        addr := ":" + strconv.Itoa(7000+i)
        Pipeline.NetworkSink(addr, Pipeline.InMemorySort(source))
        sortAddr = append(sortAddr, addr)
    }

    var sortResult []<-chan int
    for _, addr := range sortAddr {
        sortResult = append(sortResult, Pipeline.NetworkSource(addr))
    }
    return Pipeline.MergeN(sortResult...)
}

func ReadData() {
    file, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    p := Pipeline.ReaderSource(bufio.NewReader(file), -1)
    count := 0
    for v := range p {
        count++
        _, _ = fmt.Println(v)
        if count >= 10 {
            break
        }
    }
}

func WriteData() {
    file, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    p := Pipeline.RandomSource(100000000)
    writer := bufio.NewWriter(file)
    Pipeline.WriterSink(writer, p)
    _ = writer.Flush()
}

func MergeDemo() {
    p := Pipeline.Merge(
        Pipeline.InMemorySort(
            Pipeline.ArraySource(9, 5, 99, 5, 6, 89, 8, 54, 9, 56, 5, 9, 58, 9, 2, 5, 9, 56, 43, 48)),
        Pipeline.InMemorySort(
            Pipeline.ArraySource(5, 9, 2, 8, 6, 9, 5, 8, 1, 76, 3, 9, 2, 87, 8, 69, 2, 9, 6, 89, 94)))
    for v := range p {
        _, _ = fmt.Println(v)
    }
}

package MergeSort

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

var startTime time.Time

func Init() {
	startTime = time.Now()
}

func ArraySource(a ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range a {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

func ReaderSource(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int)
	go func() {
		buf := make([]byte, 8)
		bytesRead := 0
		for {
			n, err := reader.Read(buf)
			bytesRead += n
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buf))
				out <- v
			}
			if err != nil || (bytesRead >= chunkSize && chunkSize != -1) {
				break
			}
		}
		close(out)
	}()
	return out
}

func WriterSink(writer io.Writer, in <-chan int) {
	for v := range in {
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(v))
		_, _ = writer.Write(buf)
	}
}

func InMemorySort(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var a []int
		for v := range in {
			a = append(a, v)
		}
		_, _ = fmt.Println("Read done: ", time.Now().Sub(startTime))

		sort.Ints(a)
		_, _ = fmt.Println("InMemorySort done: ", time.Now().Sub(startTime))

		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
		_, _ = fmt.Println("Merge done: ", time.Now().Sub(startTime))
	}()
	return out
}

func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}

	m := len(inputs) / 2
	return Merge(
		MergeN(inputs[:m]...),
		MergeN(inputs[m:]...))
}

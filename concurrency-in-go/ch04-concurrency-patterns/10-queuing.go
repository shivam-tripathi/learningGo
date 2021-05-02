package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

// Buffered channels are a type of queue

// Queuing is one of the last techniques we employ for optimization. Adding
// queuing prematurely can hide synchronization issues such as deadlocks and
// livelocks. As our program further converges towards correctedness, we may
// find we need more or less queuing.

// - Queuing will almost never actually speed up the program, it will only make
// it to behave differently.
// - Queuing can reduce "blocking" time for a stage in pipeline, making it more responsive
// for a particular stage. True utility of queues is to decouple stages so that runtime
// of one stage has no effect on runtime of another stage.
// - Queuing can increase overall performance of the system if:
// 		* Batching requests can save time
// 		* Delays in a stage produce feedback loop into the system
func example_01(done <-chan interface{}) {
	// This pipeline contains 4 stages:
	// 1. Repeat stage that generates endless stream of 0s
	// 2. A stage that cancels the previous stage after seeing 3 zeros
	// 3. A "short" stage that sleeps 3 seconds
	// 4. A "long" stage that sleeps 4 seconds
	// When executing, first two steps finish almost immediately
	// (repeat produces zero on demand and zeros take just three of it)
	// short sleeps for 1 second, but the next stage in pipeline long sleeps for
	// full 4 seconds. So short is basically left stuck till long finished.
	// This can be an issue if incoming is fast and we might loose data if timeouts happen
	// To remedy this, we use buffer which can store 2 values
	repeatChan := repeat(done, 3)
	takeChan := take(done, 3, repeatChan)
	shortSleep := sleep(done, 1*time.Second, takeChan, "shortSleep")
	// We include buffer inbetween
	buffer := buffer(done, 2, shortSleep)
	longSleep := sleep(done, 4*time.Second, buffer, "longSleep")
	collect := []int{}
loop:
	for {
		select {
		case <-done:
			break loop
		case char, ok := <-longSleep:
			if !ok {
				fmt.Println("Stopping Long ::", char, ok)
				break loop
			}
			collect = append(collect, char)
		}
	}
	fmt.Println(collect)
}

// Batching using bufio package
func performWrite(b *testing.B, writer io.Writer) {
	done := make(chan interface{})
	defer close(done)
	b.ResetTimer()
	for bt := range take(done, b.N, repeat(done, 23)) {
		writer.Write([]byte{byte(bt)})
	}
}
func tmpFileOrFatal() *os.File {
	file, err := ioutil.TempFile("", "tmp")
	if err != nil {
		// log.Fatal("error: %v", err)
		log.Fatal(err)
	}
	return file
}

// the buffered write is faster than the unbuffered write. This is because in bufio.Writer,
// the writes are queued internally into a buffer until a sufficient chunk has been
// accumulated, and then the chunk is written out.
func benchmarkUnbufferedWrite(b *testing.B) {
	performWrite(b, tmpFileOrFatal())
}
func benchmarkBufferedWrite(b *testing.B) {
	bufferedFile := bufio.NewWriter(tmpFileOrFatal())
	performWrite(b, bufferedFile)
}

// Queuing can also help if your algorithm can be optimized by supporting lookbehinds, or ordering
// When there is a delay in a stage causes more input into the pipeline, is a little more
// difficult to spot, but also more important because it can lead to a systemic collapse
// of your upstream systems. This idea is often referred to as a negative feedback loop,
// downward-spiral, or even death-spiral.
// A recurrent relation exists between the pipeline and its upstream systems; the rate at
// which upstream stages or systems submit new requests is somehow linked to how efficient
// the pipeline is. If the efficiency of the pipeline drops below a certain critical threshold,
// the systems upstream from the pipeline begin increasing their inputs into the pipeline, which
// causes the pipeline to lose more efficiency, and the death-spiral begins. Without some sort of
// fail-safe, the system utilizing the pipeline will never recover. By introducing a queue at
// the entrance to the pipeline, you can break the feedback loop at the cost of creating lag for
// requests. From the perspective of the caller, the request appears to be processing, but at a
// very slow pace. We must make sure to handle timed out clients as well, else it may lead to dead
// message being processed again decreasing queue's efficiency.

// Queuing should be implemented at:
// 1. Entrance of the pipeline
// 2. In stages where batching will lead to higher efficiency

/* ===================================== Helpers (can ignore) ======================================= */

func repeat(done <-chan interface{}, val int) chan int {
	repeatChan := make(chan int)
	go func() {
		defer close(repeatChan)
		count := 0
	loop:
		for {
			select {
			case <-done:
				fmt.Println("Repeat ::", val, "Breaking")
				break loop
			case repeatChan <- val:
				count++
				fmt.Println("Repeat ::", val, "(", count, ")")
			}
		}
	}()
	return repeatChan
}

func take(done <-chan interface{}, times int, producer chan int) chan int {
	consumer := make(chan int)
	go func() {
		defer close(consumer)
		count := 0
	loop:
		for {
			select {
			case <-done:
				fmt.Println("Take :: Break")
				break
			case val, ok := <-producer:
				count++
				if !ok {
					fmt.Println("Take :: Producer is closed.")
					break loop
				} else {
					fmt.Println("Take ::", val, "(", count, "/", times, ")")
					consumer <- val
				}
				if count >= times {
					fmt.Println("Take ::", times, "Breaking now")
					break loop
				}
			}
		}
	}()
	return consumer
}

func sleep(done <-chan interface{}, timeMS time.Duration, inp <-chan int, namespace string) chan int {
	sleepChan := make(chan int)
	go func() {
		defer close(sleepChan)
	loop:
		for {
			select {
			case <-done:
				fmt.Println("Sleep :: Done", namespace)
				break loop
			case msg, ok := <-inp:
				if !ok {
					fmt.Println("Sleep :: Break. inp channel closed", namespace)
					break loop
				}
				fmt.Println("Sleep ::", namespace, timeMS, "Msg ::", msg)
				time.Sleep(timeMS)
				sleepChan <- msg
			}
		}
	}()
	return sleepChan
}

func buffer(done <-chan interface{}, bufferSize int, inp <-chan int) <-chan int {
	bufferedChannel := make(chan int, bufferSize)
	go func() {
		defer close(bufferedChannel)
	loop:
		for {
			select {
			case <-done:
				break loop
			case val, ok := <-inp:
				if !ok {
					break loop
				}
				fmt.Println("Buffer ::", val)
				bufferedChannel <- val
			}
		}
	}()
	return bufferedChannel
}

// func main() {
// 	done := make(chan interface{})
// 	defer close(done)
// 	example_01(done)
// 	done <- true
// }

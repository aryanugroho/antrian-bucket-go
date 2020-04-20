package antrian

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Antrian struct {
	Slot  int
	queue []int
}

func (a *Antrian) Start(process []int) {
	processed := 0
	for processed < len(process) {
		fmt.Println(processed, a.queue)
		if len(a.queue) < a.Slot {
			wg.Add(1)
			a.queue = append(a.queue, 1)
			go a.serve(process[processed])
			processed++
		}
	}
}

func (a *Antrian) serve(howLong int) {
	defer func() {
		if len(a.queue) > 0 {
			a.queue = a.queue[:len(a.queue)-1]
		}
		fmt.Println("done... ", howLong, " ", time.Now())
		wg.Done()
	}()

	fmt.Println("processing... ", howLong, " ", time.Now())
	durationString := fmt.Sprintf("%ss", strconv.Itoa(howLong))
	duration, _ := time.ParseDuration(durationString)
	time.Sleep(duration)
}

func (a *Antrian) Close() {
	fmt.Println("waiting existing process")
	wg.Wait()
}

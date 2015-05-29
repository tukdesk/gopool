package main

import (
	"log"
	"runtime"
	"time"

	"github.com/tukdesk/gopool"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	cfg := gopool.Config{
		Size:        3,
		Constructor: func() (interface{}, error) { return time.Now(), nil },
	}
	p, err := gopool.NewPool(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	ct := make(chan int)

	count := 20
	for i := 0; i < count; i++ {
		go func(num int, out chan int) {
			t, _ := p.Get()
			defer p.Put(t)

			log.Println(num, t.(time.Time).Nanosecond())

			out <- 1
		}(i, ct)
	}

	for i := 0; i < count; i++ {
		<-ct
	}

	p.Close()
}

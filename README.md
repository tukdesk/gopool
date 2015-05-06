#### gopool

##### installation
```
go get -u github.com/tukdesk/gopool
```

##### example

```
package main

import (
	"log"
	"runtime"
	"time"

	"github.com/tukdesk/gopool"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	p, err := gopool.NewPool(3, func() (interface{}, error) { return time.Now(), nil })
	if err != nil {
		log.Fatalln(err)
	}

	ct := make(chan int)

	count := 20
	for i := 0; i < count; i++ {
		go func(num int, out chan int) {
			t, _ := p.Get().(time.Time)
			defer p.Put(t)

			log.Println(num, t.Nanosecond())

			out <- 1
		}(i, ct)
	}

	for i := 0; i < count; i++ {
		<-ct
	}
}

```
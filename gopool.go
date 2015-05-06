package gopool

import (
	"fmt"
)

const (
	defaultSize = 10
)

type Pool struct {
	pool chan interface{}
}

func NewPool(cfg Config) (*Pool, error) {
	if cfg.Constructor == nil {
		return nil, fmt.Errorf("constructor required")
	}

	if cfg.Size < 1 {
		cfg.Size = defaultSize
	}

	pool := &Pool{
		pool: make(chan interface{}, cfg.Size),
	}

	for i := 0; i < cfg.Size; i++ {
		element, err := cfg.Constructor()
		if err != nil {
			return nil, err
		}
		pool.Put(element)
	}
	return pool, nil
}

func (this *Pool) Get() interface{} {
	return <-this.pool
}

func (this *Pool) Put(x interface{}) {
	select {
	case this.pool <- x:

	default:
		// in case the pool is full
	}
}

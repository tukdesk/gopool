package gopool

import (
	"fmt"
)

var ErrPoolClosed = fmt.Errorf("pool closed")

const (
	defaultMin  = 1
	defaultSize = 10
)

type Pool struct {
	cfg    Config
	pool   chan interface{}
	closed bool
}

func NewPool(cfg Config) (*Pool, error) {
	if cfg.Constructor == nil {
		return nil, fmt.Errorf("constructor required")
	}

	if cfg.Size < 1 {
		cfg.Size = defaultSize
	}

	if cfg.Min < 0 || cfg.Min > cfg.Size {
		cfg.Min = defaultMin
	}

	pool := &Pool{
		cfg:  cfg,
		pool: make(chan interface{}, cfg.Size),
	}

	for i := 0; i < cfg.Min; i++ {
		element, err := cfg.Constructor()
		if err != nil {
			return nil, err
		}
		pool.Put(element)
	}
	return pool, nil
}

func (this *Pool) Get() (interface{}, error) {
	if this.closed {
		return nil, ErrPoolClosed
	}

	select {
	case x := <-this.pool:
		return x, nil
	default:
		return this.cfg.Constructor()
	}
}

func (this *Pool) Put(x interface{}) {
	if this.closed {
		return
	}

	select {
	case this.pool <- x:

	default:
		// in case the pool is full
	}
}

func (this *Pool) Stop() chan interface{} {
	this.closed = true
	return this.pool
}

func (this *Pool) Close() {
	this.Stop()
	close(this.pool)
	return
}

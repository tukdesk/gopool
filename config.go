package gopool

type Config struct {
	Min         int
	Size        int
	Constructor func() (interface{}, error)
}

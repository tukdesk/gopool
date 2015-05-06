package gopool

type Config struct {
	Size        int
	Constructor func() (interface{}, error)
}

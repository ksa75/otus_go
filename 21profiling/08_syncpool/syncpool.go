package syncpool

import (
	"strings"
	"sync"
)

var (
	pool = sync.Pool{
		New: func() any {
			return &strings.Builder{}
		},
	}
)

func Slow() string {
	builder := strings.Builder{}
	builder.WriteString("Hello")
	return builder.String()
}

func Fast() string {
	builder := pool.Get().(*strings.Builder)
	defer pool.Put(builder)
	builder.Reset()
	builder.WriteString("Hello")
	res := builder.String()
	return res
}

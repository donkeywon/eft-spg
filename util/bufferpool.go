package util

import (
	"bytes"
	"sync"
)

var (
	_bufferPool = &sync.Pool{New: func() interface{} {
		return &Buffer{
			Buffer: &bytes.Buffer{},
		}
	}}
)

type Buffer struct {
	*bytes.Buffer
	p *sync.Pool
}

func GetBuffer() *Buffer {
	b := _bufferPool.Get().(*Buffer)
	b.p = _bufferPool
	b.Reset()
	return b
}

func (b *Buffer) Free() {
	b.p.Put(b)
}

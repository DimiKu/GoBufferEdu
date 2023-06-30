package GoBufferEdu

import "time"

package main

import (
"bytes"
"go/ast"
"io"
"log"
"strings"
"time"
)


type Buffer struct {
	buff []Data
	createTime time.Time
	ttl time.Ticker
}

type Data struct {
	data []byte
	createTime time.Time
	ttl time.Ticker
}

func (b *Buffer)AddDataInBuffer(data Data, ttl time.Ticker) error {
	b.buff = append(b.buff, data)
	b.createTime = time.Now()
	b.ttl = ttl
	return nil
}

go func CheckBuffer(buff []Data) []byte {
	for i := range buff {
		if buff[i].ttl
	}
}

func main() {

}
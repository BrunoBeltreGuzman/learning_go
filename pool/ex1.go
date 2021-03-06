package main

import (
	"bytes"
	"sync"
)

var pool = sync.Pool{
	// New creates an object when the pool has nothing available to return.
	// New must return an interface{} to make it flexible. You have to cast
	// your type after getting it.
	New: func() interface{} {
		// Pools often contain things like *bytes.Buffer, which are
		// temporary and re-usable.
		return &bytes.Buffer{}
	},
}

func main() {

	//poo()

}

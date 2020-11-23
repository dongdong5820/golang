package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := []interface{}{"a", "b", "c", "d", "e"}
	fmt.Println(ArrayRand(a))
}

//
func ArrayRand(elements []interface{}) []interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := make([]interface{}, len(elements))
	for i, v := range r.Perm(len(elements)) {
		n[i] = elements[v]
	}

	return n
}

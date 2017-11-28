package robotname

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

const (
	testVersion = 1
)

type Robot struct {
	Id   []byte
	Used int
}

//used 0 is initial case where Id is still empty
func (r *Robot) Name() string {
	if r.Used == 0 {
		r.Reset()
	}
	return fmt.Sprintf("%s", r.Id)
}

// generate name. when different than the original, return
func (r *Robot) Reset() {
	prev := r.Id
	n := GenerateNew()
	for bytes.Equal(prev, n) {
		n = GenerateNew()
	}
	r.Id = n
	r.Used = 1

}

// 2 capital letters and 3 numbers generated randomly
func GenerateNew() []byte {
	rand.Seed(time.Now().UTC().UnixNano())
	return []byte{byte(rand.Intn(90-65) + 65), byte(rand.Intn(90-65) + 65), byte(rand.Intn(57-48) + 48), byte(rand.Intn(57-48) + 48), byte(rand.Intn(57-48) + 48)}
}

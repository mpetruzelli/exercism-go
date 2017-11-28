package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))
var four = big.NewInt(4)

func PrivateKey(i *big.Int) *big.Int {
	p := new(big.Int).Sub(i, four)
	return p.Add(four, p.Rand(r, p))

}

func PublicKey(i, j *big.Int, k int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(k), i, j)

}

func SecretKey(i, j, k *big.Int) *big.Int {
	return new(big.Int).Exp(j, i, k)
}

func NewPair(i *big.Int, k int64) (*big.Int, *big.Int) {
	f := PrivateKey(i)
	return f, PublicKey(f, i, k)
}

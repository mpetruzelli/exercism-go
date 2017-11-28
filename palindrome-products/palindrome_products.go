package palindrome

import (
	"errors"
)

const (
	testVersion = 1
)

type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

func Products(fmin, fmax int) (Product, Product, error) {

	if fmax < fmin {
		return Product{}, Product{}, errors.New("fmin > fmax...")
	}
	var palindromes []Product
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			prod := i * j
			if IsPalindrom(prod) {
				p := Product{prod, [][2]int{{i, j}}}
				palindromes = Insert(palindromes, p)
			}
		}
	}
	if len(palindromes) <= 0 {
		return Product{}, Product{}, errors.New("no palindromes")
	}
	return palindromes[0], palindromes[len(palindromes)-1], nil
}

func IsPalindrom(num int) bool {
	n := num
	rev := 0
	for num > 0 {
		dig := num % 10
		rev = rev*10 + dig
		num = num / 10
	}
	return n == rev
}

func Insert(prods []Product, p Product) []Product {
	for i := 0; i < len(prods); i++ {
		if p.Product < prods[i].Product {
			return append(prods[:i], append([]Product{p}, prods[i:]...)...)
		} else if prods[i].Product == p.Product {
			prods[i].Factorizations = append(prods[i].Factorizations, p.Factorizations...)
			return prods
		}
	}
	return append(prods, p)
}

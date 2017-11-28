package perfect

import (
	"errors"
)

type Classification string

const (
	testVersion                            = 1
	ClassificationDeficient Classification = "ClassificationDeficient"
	ClassificationAbundant  Classification = "ClassificationAbundant"
	ClassificationPerfect   Classification = "ClassificationPerfect"
)

var ErrOnlyPositive = errors.New("only positive")

func Classify(n uint64) (Classification, error) {
	if n < 1 {
		return "", ErrOnlyPositive
	}
	sum := Aliquot(int(n))
	if sum == int(n) {
		return ClassificationPerfect, nil
	}
	if sum > int(n) {
		return ClassificationAbundant, nil
	}
	return ClassificationDeficient, nil
}

func Aliquot(n int) (sum int) {
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return
}

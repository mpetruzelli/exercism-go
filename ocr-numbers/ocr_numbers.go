package ocr

import (
	"fmt"
	"strings"
)

var digit = map[string]string{
	"_ | ||_|": "0",
	"|  |":     "1",
	"_  _||_":  "2",
	"_  _| _|": "3",
	"|_|  |":   "4",
	"_ |_  _|": "5",
	"_ |_ |_|": "6",
	"_   |  |": "7",
	"_ |_||_|": "8",
	"_ |_| _|": "9",
}

func recognizeDigit() {
}

func Recognize(r string) []string {

	fmt.Println(strings.Join(strings.Split(strings.TrimSpace(r), "\n")[:], ""))
	return []string{digit[strings.Join(strings.Split(strings.TrimSpace(r), "\n")[:], "")]}
}

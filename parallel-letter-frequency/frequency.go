package letter

const (
	testVersion = 1
)

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(i []string) FreqMap {
	ch := make(chan FreqMap, len(i)) //create buffered channel
	o := FreqMap{}
	for _, s := range i {
		go func(s string) {
			ch <- Frequency(s)
		}(s)
	}
	for range i {
		for r, c := range <-ch {
			o[r] += c
		}
	}
	return o
}

/** other solution
func ConcurrentFrequency(listOfSentences []string) FreqMap {

	ch := make(chan FreqMap)

	f := func(sentence string) {
		ch <- Frequency(sentence)
	}

	for _, sentence := range listOfSentences {
		go f(sentence)
	}

	m := FreqMap{}

	for range listOfSentences {
		for r, n := range <-ch {
			m[r] += n
		}
	}

	return m
}*/

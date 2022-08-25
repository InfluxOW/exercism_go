package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(strings []string) FreqMap {
	m := make(FreqMap)
	ch := make(chan FreqMap)
	var wg sync.WaitGroup

	wg.Add(len(strings))

	for _, s := range strings {
		go func(s string) {
			ch <- Frequency(s)
			wg.Done()
		}(s)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for fm := range ch {
		for r, f := range fm {
			m[r] += f
		}
	}

	return m
}

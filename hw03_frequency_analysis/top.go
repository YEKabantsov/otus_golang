package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type KeyValue struct {
	Key   string
	Value int
}

func Top10(str string) []string {
	if len(str) == 0 {
		return make([]string, 0)
	}

	words := strings.Fields(str)
	topWords := make(map[string]int)

	for _, word := range words {
		topWords[word]++
	}

	s := make([]KeyValue, 0, len(topWords))
	for k, v := range topWords {
		s = append(s, KeyValue{k, v})
	}

	sort.Slice(s, func(i, j int) bool {
		if s[i].Value > s[j].Value {
			return true
		}
		if s[i].Value < s[j].Value {
			return false
		}
		return s[i].Key < s[j].Key
	})

	res := make([]string, 0, 10)
	iterations := 10

	if len(s) < 10 {
		iterations = len(s)
	}

	for i := 0; i < iterations; i++ {
		res = append(res, s[i].Key)
	}

	return res
}

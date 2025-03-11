package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	l string
	r string
}

type Token struct {
	pair Pair
}

func main() {

	freq := map[Token]int{}

	text := "The original BPE algorithm operates by iteratively replacing the most common contiguous sequences of characters in a target text with unused 'placeholder' bytes. The iteration ends when no sequences can be found, leaving the target text effectively compressed. Decompression can be performed by reversing this process, querying known placeholder terms against their corresponding denoted sequence, using a lookup table. In the original paper, this lookup table is encoded and stored alongside the compressed text."

	for i := 0; i < len(text)-1; i++ {
		token := Token{
			pair: Pair{
				l: string(text[i]),
				r: string(text[i+1]),
			},
		}

		_, exists := freq[token]

		if !exists {
			freq[token] = 1
		} else {
			freq[token] = freq[token] + 1
		}
	}

	tokenSort(freq, 10)

}

func tokenSort(freq map[Token]int, topK int) {
	count := 0
	pairs := make([]Token, 0, len(freq))

	for key := range freq {
		pairs = append(pairs, key)
	}

	sort.Slice(pairs, func(i, j int) bool { return freq[pairs[i]] > freq[pairs[j]] })

	for _, k := range pairs {

		if count >= topK {
			return
		}

		fmt.Printf("(%s %s) -> %d \n", k.pair.l, k.pair.r, freq[k])

		count++
	}

}

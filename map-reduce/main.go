package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"
)

type pair struct {
	word string
	freq int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go <input_file>")
		os.Exit(1)
	}
	inputFile := os.Args[1]
	stopFile := "stop_words.txt"

	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	chunks := partition(string(data), 200)
	start := time.Now()
	pairsLists := mapPhase(split, chunks)
	filteredPairs := removeStopWords(stopFile, pairsLists)
	wordFreqs := reducePhase(countWords, filteredPairs)

	fmt.Printf("Execution time: %v\n\n", time.Since(start))

	sort.Slice(wordFreqs, func(i, j int) bool {
		return wordFreqs[i].freq > wordFreqs[j].freq
	})

	limit := min(25, len(wordFreqs))
	for i := 0; i < limit; i++ {
		fmt.Printf("%s - %d\n", wordFreqs[i].word, wordFreqs[i].freq)
	}
}

func partition(data string, nlines int) []string {
	lines := strings.Split(data, "\n")
	numChunks := int(math.Ceil(float64(len(lines)) / float64(nlines)))
	chunks := []string{}
	for i := 0; i < numChunks; i++ {
		start := i * nlines
		end := min((i+1)*nlines, len(lines))
		chunk := strings.Join(lines[start:end], " ")
		chunks = append(chunks, chunk)
	}
	return chunks
}

func mapPhase(splitFunc func(string) []pair, dataChunks []string) [][]pair {
	pairsLists := [][]pair{}
	for _, chunk := range dataChunks {
		pairs := splitFunc(chunk)
		pairsLists = append(pairsLists, pairs)
	}
	return pairsLists
}

func split(text string) []pair {
	normalized := normalize(text)
	re := regexp.MustCompile(`\s+`)
	words := re.Split(normalized, -1)
	pairs := []pair{}
	for _, w := range words {
		if w != "" {
			pairs = append(pairs, pair{word: w, freq: 1})
		}
	}
	return pairs
}

func removeStopWords(stopFile string, pairsLists [][]pair) [][]pair {
	stopwords := getStopWords(stopFile)
	filteredLists := [][]pair{}
	for _, pairs := range pairsLists {
		filtered := []pair{}
		for _, p := range pairs {
			if _, found := stopwords[p.word]; !found {
				filtered = append(filtered, p)
			}
		}
		filteredLists = append(filteredLists, filtered)
	}
	return filteredLists
}

func reducePhase(reduceFunc func([]pair, []pair) []pair, pairsLists [][]pair) []pair {
	n := len(pairsLists)
	if n == 0 {
		return nil
	}
	for n > 1 {
		newLists := [][]pair{}
		for i := 0; i < n; i += 2 {
			if i+1 < n {
				merged := reduceFunc(pairsLists[i], pairsLists[i+1])
				newLists = append(newLists, merged)
			} else {
				newLists = append(newLists, pairsLists[i])
			}
		}
		pairsLists = newLists
		n = len(pairsLists)
	}
	return pairsLists[0]
}

func countWords(pairs1, pairs2 []pair) []pair {
	freqMap := map[string]int{}
	for _, p := range pairs1 {
		freqMap[p.word] += p.freq
	}
	for _, p := range pairs2 {
		freqMap[p.word] += p.freq
	}
	result := []pair{}
	for w, f := range freqMap {
		result = append(result, pair{word: w, freq: f})
	}
	return result
}

func getStopWords(file string) map[string]struct{} {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading stop words file: %v", err)
	}
	stopwords := map[string]struct{}{}
	for _, sw := range strings.Split(string(data), ",") {
		sw = strings.TrimSpace(strings.ToLower(sw))
		if sw != "" {
			stopwords[sw] = struct{}{}
		}
	}
	for r := 'a'; r <= 'z'; r++ {
		stopwords[string(r)] = struct{}{}
	}
	return stopwords
}

func normalize(text string) string {
	clean := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return unicode.ToLower(r)
		}
		return ' '
	}, text)
	return clean
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go <path-to-input-file>")
		os.Exit(1)
	}

	stopFile, err := os.Open("monolithic/stop_words.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening ../stop_words.txt: %v\n", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(stopFile)
	scanner.Split(bufio.ScanLines)
	stopWords := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		for _, p := range parts {
			p = strings.TrimSpace(strings.ToLower(p))
			if p != "" {
				stopWords = append(stopWords, p)
			}
		}
	}
	stopFile.Close()

	for ch := 'a'; ch <= 'z'; ch++ {
		stopWords = append(stopWords, string(ch))
	}

	inFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer inFile.Close()

	words := []string{}
	freqs := []int{}

	inScanner := bufio.NewScanner(inFile)
	for inScanner.Scan() {
		line := inScanner.Text()
		runes := []rune(line)
		start := -1
		for i, r := range runes {
			if start == -1 {
				if unicode.IsLetter(r) || unicode.IsDigit(r) {
					start = i
				}
			} else {
				if !(unicode.IsLetter(r) || unicode.IsDigit(r)) {
					word := strings.ToLower(string(runes[start:i]))
					isStop := false
					for _, sw := range stopWords {
						if word == sw {
							isStop = true
							break
						}
					}
					if !isStop && word != "" {
						found := false
						idx := 0
						for j, w := range words {
							if w == word {
								freqs[j]++
								found = true
								idx = j
								break
							}
						}
						if !found {
							words = append(words, word)
							freqs = append(freqs, 1)
							idx = len(words) - 1
						}
						for k := idx - 1; k >= 0; k-- {
							if freqs[idx] > freqs[k] {
								words[k], words[idx] = words[idx], words[k]
								freqs[k], freqs[idx] = freqs[idx], freqs[k]
								idx = k
							} else {
								break
							}
						}
					}
					start = -1
				}
			}
		}
		if start != -1 {
			word := strings.ToLower(string(runes[start:]))
			isStop := false
			for _, sw := range stopWords {
				if word == sw {
					isStop = true
					break
				}
			}
			if !isStop && word != "" {
				found := false
				idx := 0
				for j, w := range words {
					if w == word {
						freqs[j]++
						found = true
						idx = j
						break
					}
				}
				if !found {
					words = append(words, word)
					freqs = append(freqs, 1)
					idx = len(words) - 1
				}
				for k := idx - 1; k >= 0; k-- {
					if freqs[idx] > freqs[k] {
						words[k], words[idx] = words[idx], words[k]
						freqs[k], freqs[idx] = freqs[idx], freqs[k]
						idx = k
					} else {
						break
					}
				}
			}
		}
	}

	limit := 25
	if len(words) < limit {
		limit = len(words)
	}
	for i := 0; i < limit; i++ {
		fmt.Printf("%s - %d\n", words[i], freqs[i])
	}
}

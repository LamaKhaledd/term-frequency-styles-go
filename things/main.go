package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

type DataStorageManager struct {
	data string
}

func NewDataStorageManager(path string) *DataStorageManager {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile(`[\W_]+`)
	cleaned := re.ReplaceAllString(string(content), " ")
	cleaned = strings.ToLower(cleaned)
	return &DataStorageManager{data: cleaned}
}

func (dsm *DataStorageManager) Words() []string {
	return strings.Fields(dsm.data)
}

type StopWordManager struct {
	stopWords map[string]bool
}

func NewStopWordManager() *StopWordManager {
	file, err := os.Open("stop_words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	stopWords := strings.Split(scanner.Text(), ",")
	for ch := 'a'; ch <= 'z'; ch++ {
		stopWords = append(stopWords, string(ch))
	}

	stopSet := make(map[string]bool)
	for _, w := range stopWords {
		stopSet[w] = true
	}

	return &StopWordManager{stopWords: stopSet}
}

func (swm *StopWordManager) IsStopWord(word string) bool {
	return swm.stopWords[word]
}

type WordFrequencyManager struct {
	wordFreqs map[string]int
}

func NewWordFrequencyManager() *WordFrequencyManager {
	return &WordFrequencyManager{wordFreqs: make(map[string]int)}
}

func (wfm *WordFrequencyManager) IncrementCount(word string) {
	wfm.wordFreqs[word]++
}

func (wfm *WordFrequencyManager) Sorted() [][2]string {
	type kv struct {
		Key   string
		Value int
	}
	var pairs []kv
	for k, v := range wfm.wordFreqs {
		pairs = append(pairs, kv{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	var result [][2]string
	for _, p := range pairs {
		result = append(result, [2]string{p.Key, fmt.Sprintf("%d", p.Value)})
	}
	return result
}

type WordFrequencyController struct {
	storageManager  *DataStorageManager
	stopWordManager *StopWordManager
	wordFreqManager *WordFrequencyManager
}

func NewWordFrequencyController(path string) *WordFrequencyController {
	return &WordFrequencyController{
		storageManager:  NewDataStorageManager(path),
		stopWordManager: NewStopWordManager(),
		wordFreqManager: NewWordFrequencyManager(),
	}
}

func (wfc *WordFrequencyController) Run() {
	for _, w := range wfc.storageManager.Words() {
		if !wfc.stopWordManager.IsStopWord(w) {
			wfc.wordFreqManager.IncrementCount(w)
		}
	}

	wordFreqs := wfc.wordFreqManager.Sorted()
	for i, p := range wordFreqs {
		if i >= 25 {
			break
		}
		fmt.Printf("%s - %s\n", p[0], p[1])
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go")
		os.Exit(1)
	}
	controller := NewWordFrequencyController(os.Args[1])
	controller.Run()
}

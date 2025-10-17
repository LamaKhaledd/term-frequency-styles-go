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

func readFile(path string) string {
    data, err := os.ReadFile(path)
    if err != nil {
        log.Fatal(err)
    }
    return string(data)
}

func filterCharsAndNormalize(data string) string {
    re := regexp.MustCompile(`[\W_]+`)
    cleaned := re.ReplaceAllString(data, " ")
    return strings.ToLower(cleaned)
}

func scan(data string) []string {
    return strings.Fields(data)
}

func removeStopWords(words []string) []string {
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

    var filtered []string
    for _, w := range words {
        if !stopSet[w] {
            filtered = append(filtered, w)
        }
    }
    return filtered
}

func frequencies(words []string) map[string]int {
    freq := make(map[string]int)
    for _, w := range words {
        freq[w]++
    }
    return freq
}

func sortFreqs(freq map[string]int) [][2]string {
    type kv struct {
        Key   string
        Value int
    }

    var pairs []kv
    for k, v := range freq {
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

func printAll(pairs [][2]string) {
    for i, p := range pairs {
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

    printAll(
        sortFreqs(
            frequencies(
                removeStopWords(
                    scan(
                        filterCharsAndNormalize(
                            readFile(os.Args[1]),
                        ),
                    ),
                ),
            ),
        ),
    )
}

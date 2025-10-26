package main

import (
    "fmt"
    "log"
    "os"
    "regexp"
    "sort"
    "strings"
)

type TFQuarantine struct {
    funcs []func(any) any
}

func (q *TFQuarantine) Bind(f func(any) any) *TFQuarantine {
    q.funcs = append(q.funcs, f)
    return q
}

func (q *TFQuarantine) Execute() {
    var value any = func() any { return nil }

    guardCallable := func(v any) any {
        if fn, ok := v.(func() any); ok {
            return fn()
        }
        return v
    }

    for _, f := range q.funcs {
        value = f(guardCallable(value))
    }

    if resultFunc, ok := value.(func() any); ok {
        resultFunc()
    } else {
        fmt.Println(value)
    }
}

func getInput() func() any {
    return func() any {
        return os.Args[1]
    }
}

func extractWords(v any) any {
    return func() any {
        path := v.(string)
        data, err := os.ReadFile(path)
        if err != nil {
            log.Fatal(err)
        }

        re := regexp.MustCompile(`[\W_]+`)
        words := strings.Fields(strings.ToLower(re.ReplaceAllString(string(data), " ")))
        return words
    }
}

func removeStopWords(v any) any {
    return func() any {
        words := v.([]string)
        stopData, err := os.ReadFile("../stop_words.txt")
        if err != nil {
            log.Fatal(err)
        }

        stopWords := strings.Split(string(stopData), ",")
        for _, r := range "abcdefghijklmnopqrstuvwxyz" {
            stopWords = append(stopWords, string(r))
        }

        stopSet := make(map[string]bool)
        for _, w := range stopWords {
            stopSet[w] = true
        }

        var result []string
        for _, w := range words {
            if !stopSet[w] {
                result = append(result, w)
            }
        }
        return result
    }
}

func frequencies(v any) any {
    words := v.([]string)
    freqs := make(map[string]int)
    for _, w := range words {
        freqs[w]++
    }
    return freqs
}

type kv struct {
    Key   string
    Value int
}

func sortFreqs(v any) any {
    freqs := v.(map[string]int)
    var pairs []kv
    for k, v := range freqs {
        pairs = append(pairs, kv{k, v})
    }
    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i].Value > pairs[j].Value
    })
    return pairs
}

func top25Freqs(v any) any {
    pairs := v.([]kv)
    return func() any {
        for i, kv := range pairs {
            if i >= 25 {
                break
            }
            fmt.Printf("%s - %d\n", kv.Key, kv.Value)
        }
        return nil
    }
}

func main() {
    (&TFQuarantine{funcs: []func(any) any{func(_ any) any { return getInput() }}}).
        Bind(extractWords).
        Bind(removeStopWords).
        Bind(frequencies).
        Bind(sortFreqs).
        Bind(top25Freqs).
        Execute()
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
)

type Message struct {
	Tag  string
	Data interface{}
}

type Actor struct {
	Inbox chan Message
}

func Send(actor *Actor, msg Message) {
	actor.Inbox <- msg
}

func DataStorageManager(path string, stopWordMgr *Actor, controller *Actor, wg *sync.WaitGroup) *Actor {
	actor := &Actor{Inbox: make(chan Message)}
	go func() {
		var data string
		for msg := range actor.Inbox {
			switch msg.Tag {
			case "init":
				file, _ := os.Open(path)
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					data += scanner.Text() + " "
				}
				file.Close()
				re := regexp.MustCompile(`[\W_]+`)
				data = strings.ToLower(re.ReplaceAllString(data, " "))
			case "send_words":
				words := strings.Fields(data)
				for _, w := range words {
					Send(stopWordMgr, Message{"filter", w})
				}
				Send(stopWordMgr, Message{"top25", controller})
			case "die":
				close(actor.Inbox)
				wg.Done()
				return
			}
		}
	}()
	return actor
}

func StopWordManager(wordFreqMgr *Actor, wg *sync.WaitGroup) *Actor {
	actor := &Actor{Inbox: make(chan Message)}
	go func() {
		stopWords := make(map[string]struct{})
		file, _ := os.Open("../stop_words.txt")
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			for _, w := range strings.Split(scanner.Text(), ",") {
				stopWords[w] = struct{}{}
			}
		}
		file.Close()
		for c := 'a'; c <= 'z'; c++ {
			stopWords[string(c)] = struct{}{}
		}

		for msg := range actor.Inbox {
			switch msg.Tag {
			case "filter":
				word := msg.Data.(string)
				if _, exists := stopWords[word]; !exists {
					Send(wordFreqMgr, Message{"word", word})
				}
			case "top25":
				Send(wordFreqMgr, Message{"top25", msg.Data})
			case "die":
				close(actor.Inbox)
				wg.Done()
				return
			}
		}
	}()
	return actor
}

func WordFrequencyManager(wg *sync.WaitGroup) *Actor {
	actor := &Actor{Inbox: make(chan Message)}
	go func() {
		freqs := make(map[string]int)
		for msg := range actor.Inbox {
			switch msg.Tag {
			case "word":
				word := msg.Data.(string)
				freqs[word]++
			case "top25":
				recipient := msg.Data.(*Actor)
				type kv struct {
					Key   string
					Value int
				}
				var freqList []kv
				for k, v := range freqs {
					freqList = append(freqList, kv{k, v})
				}
				sort.Slice(freqList, func(i, j int) bool {
					return freqList[i].Value > freqList[j].Value
				})
				Send(recipient, Message{"top25", freqList})
			case "die":
				close(actor.Inbox)
				wg.Done()
				return
			}
		}
	}()
	return actor
}

func WordFrequencyController(wg *sync.WaitGroup) *Actor {
	actor := &Actor{Inbox: make(chan Message)}
	go func() {
		for msg := range actor.Inbox {
			switch msg.Tag {
			case "top25":
				freqList := msg.Data.([]struct {
					Key   string
					Value int
				})
				for i, kv := range freqList {
					if i >= 25 {
						break
					}
					fmt.Println(kv.Key, "-", kv.Value)
				}
				wg.Done()
				return
			}
		}
	}()
	return actor
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file>")
		return
	}
	path := os.Args[1]
	var wg sync.WaitGroup
	wg.Add(4)

	wordFreqMgr := WordFrequencyManager(&wg)
	stopWordMgr := StopWordManager(wordFreqMgr, &wg)
	controller := WordFrequencyController(&wg)
	storageMgr := DataStorageManager(path, stopWordMgr, controller, &wg)

	Send(storageMgr, Message{"send_words", nil})

	wg.Wait()
}

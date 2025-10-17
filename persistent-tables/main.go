package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"persistent-tables/db"
	"regexp"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func extractWords(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile(`[\W_]+`)
	cleaned := re.ReplaceAllString(string(content), " ")
	cleaned = strings.ToLower(cleaned)
	words := strings.Fields(cleaned)

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

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file path")
	}
	filePath := os.Args[1]

	conn, err := sql.Open("sqlite3", "tf.db")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	queries := db.New(conn)
	ctx := context.Background()

	document, err := queries.GetDocumentByName(ctx, filePath)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	if document.ID == 0 {
		if err := queries.InsertDocument(ctx, filePath); err != nil {
			log.Fatal(err)
		}
		document, err = queries.GetDocumentByName(ctx, filePath)
		if err != nil {
			log.Fatal(err)
		}
	}

	words := extractWords(filePath)

	var maxWordID sql.NullInt64
	row := conn.QueryRow("SELECT MAX(id) FROM words")
	if err := row.Scan(&maxWordID); err != nil {
		log.Fatal(err)
	}
	wordID := int64(0)
	if maxWordID.Valid {
		wordID = maxWordID.Int64 + 1
	}

	charID := int64(0)
	for _, w := range words {
		word := db.InsertWordParams{
			ID:    sql.NullInt64{Int64: wordID, Valid: true},
			DocID: document.ID,
			Value: w,
		}
		if err := queries.InsertWord(ctx, word); err != nil {
			log.Fatal(err)
		}

		for _, c := range w {
			char := db.InsertCharacterParams{
				ID:     sql.NullInt64{Int64: charID, Valid: true},
				WordID: wordID,
				Value:  string(c),
			}
			if err := queries.InsertCharacter(ctx, char); err != nil {
				log.Fatal(err)
			}
			charID++
		}
		wordID++
	}

	topWords, err := queries.TopWords(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, w := range topWords {
		fmt.Printf("%s - %d\n", w.Value, w.Count)
	}
}


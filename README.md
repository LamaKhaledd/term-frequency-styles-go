<a name="readme-top"></a>
<div align="center">
  <img src="third.PNG" alt="Project banner image">
  <br>
  <h1>Exercises in Programming Style – Go Implementations</h1>
  <strong>Exploring multiple programming paradigms in Go — one problem, many perspectives 💡</strong>
  <br><br>
  <img src="fourth.PNG" alt="Supporting image">
</div>

<br>

<div align="center">
  <p>
    Re-implementing the term-frequency task from Cristina Videira Lopes’ <em>Exercises in Programming Style</em> using Go.  
    This project demonstrates how constraints and idioms shape design choices across diverse programming styles.
  </p>
</div>

<br>

<details>
  <summary><h2>📚 Table of Contents</h2></summary>
  <ol>
    <li><a href="#intro">Introduction (What’s this project?)</a></li>
    <li><a href="#goal">Goal & Motivation</a></li>
    <li><a href="#chapters">Chapters & Styles Covered</a></li>
    <li><a href="#structure">Folder Structure</a></li>
    <li><a href="#techstack">Built With</a></li>
    <li><a href="#getting-started">Getting Started</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contribution">Contribution</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<br>

<a name="intro"></a>
## 💡 Introduction

This project explores how the same computational task — **term frequency analysis** — can be expressed in multiple programming styles using the **Go programming language**.

Inspired by the book  
> **_Exercises in Programming Style_**  
> by Cristina Videira Lopes  

Each chapter in the book introduces a distinct style defined by **constraints, design principles, and idioms**.  
By re-implementing them in Go, this project reveals how Go’s type system, concurrency model, and simplicity interact with various paradigms.

<br>

<a name="goal"></a>
## 🎯 Goal & Motivation

The purpose of this project is to:

- Deepen understanding of software design and architectural trade-offs.  
- Observe how Go’s language features affect different programming paradigms.  
- Practice writing idiomatic Go under unusual or restricted design constraints.  
- Reflect on code readability, maintainability, and performance across styles.

Each implementation solves the **same task**:

> Count the frequency of each word in a text file (excluding stop words) and display the top 25 most frequent words.

<br>

<a name="chapters"></a>
## 📘 Chapters & Styles Covered

| Chapter | Style | Description |
|----------|--------|-------------|
| Preface & Prologue | — | Conceptual foundation and motivation for stylistic constraints. |
| 1 | **Monolithic** | A single large function — minimal abstraction, everything inline. |
| 2 | **Pipeline** | A data-flow of transformations; no shared state. |
| 3 | **Things** | Encapsulation through “objects” — entities sending messages. |
| 4 | **Persistent Tables** | Data-driven approach using key/value tables (maps or databases). |
| 5 | **Quarantine** | Pure vs impure code — isolate I/O from computation. |
| 6 | **Actors** | Concurrency via message-passing between independent processes. |
| 7 | **Map Reduce** | Divide the problem into mappers and reducers for parallel processing. |

Optional extensions:
- Add styles from later chapters.
- Include tests, database support, and CI/CD pipelines.

<br>

<a name="structure"></a>
## 🗂 Folder Structure

```
exercises-in-style-go/
├── monolithic/
│   ├── main.go
│   └── README.md
├── pipeline/
│   ├── main.go
│   └── README.md
├── things/
│   ├── main.go
│   └── README.md
├── persistent_tables/
│   ├── main.go
│   └── README.md
├── quarantine/
│   ├── main.go
│   └── README.md
├── actors/
│   ├── main.go
│   └── README.md
├── map_reduce/
│   ├── main.go
│   └── README.md
├── db/
│   ├── migrations/       # Goose migration files
│   ├── queries/          # SQLC generated queries
│   ├── schema.sql
│   └── README.md
├── .github/
│   └── workflows/ci.yml  # GitHub Actions CI/CD pipeline
└── README.md  ← this file
```

Each chapter directory includes:
- `main.go` — the Go implementation  
- `README.md` — summary of constraints & reflection on Go implementation  
- For database support: `sqlc` and `goose` manage schema and queries  

<p align="right">(<a href="#readme-top">⬆️ Back to top</a>)</p>
<br>

<a name="techstack"></a>
## 🛠 Built With

* [![Go][GoBadge]][GoURL] – Open-source programming language designed for simplicity, concurrency, and efficiency.  
* [![SQLC][SQLCBadge]][SQLCURL] – Generate type-safe Go code from SQL queries.  
* [![Goose][GooseBadge]][GooseURL] – Database migration management for Go projects.  
* [![GitHub Actions][ActionsBadge]][ActionsURL] – CI/CD automation for builds and tests.  
* [![GitHub][GitHubBadge]][GitHubURL] – Version control and collaboration.  
* [![VSCode][VSCodeBadge]][VSCodeURL] – Code editor for development.  

<br>

<a name="getting-started"></a>
## 🚀 Getting Started

### Prerequisites

Ensure you have installed:

- **Go** (≥ 1.20)  
- **Git**  
- **Goose CLI** (for migrations)  
- **SQLC** (for generating query code)  

### Clone the Repository

```bash
git clone https://github.com/YourUserName/exercises-in-style-go.git
cd exercises-in-style-go
```

### Run Migrations

```bash
goose -dir ./db/migrations postgres "your-dsn" up
```

### Generate SQLC Code

```bash
sqlc generate
```

### Run an Implementation

```bash
cd monolithic
go run main.go ../data/input.txt ../data/stop_words.txt
```

### Run Tests

```bash
go test ./...
```

### Build

```bash
go build -o monolithic/main monolithic/main.go
```

<br>

<a name="usage"></a>
## 💡 Usage

Each program expects:
1. A path to a text file to analyze.  
2. A path to a stop-words file (or built-in list).  

Output: top 25 most frequent words in the text.

The CI/CD pipeline (via GitHub Actions) ensures:
- Code builds on push and pull requests.  
- Unit tests run automatically.  
- Database migrations are validated.  

<br>

<a name="contribution"></a>
## 🤝 Contribution

Contributions are welcome!  
You can:
- Add more styles from later chapters  
- Write unit tests or benchmarks  
- Improve database schema or queries  
- Enhance the CI/CD pipeline  
- Improve documentation or examples  

To contribute:
1. Fork the repo  
2. Create a feature branch (`git checkout -b feature/NewStyle`)  
3. Commit your changes (`git commit -m "Add new style"`)  
4. Push to the branch (`git push origin feature/NewStyle`)  
5. Create a Pull Request  

<p align="right">(<a href="#readme-top">⬆️ Back to top</a>)</p>
<br>

<a name="contact"></a>
## 📬 Contact

**Author:** Lama Ibrahim  
📧 **Email:** [lama.ibrahim@outlook.sa](mailto:lama.ibrahim@outlook.sa)  
🔗 **GitHub:** [@LamaKhaledd](https://github.com/LamaKhaledd)

<br><br>

<!-- MARKDOWN LINKS & IMAGES -->
[GoBadge]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[GoURL]: https://go.dev/
[GitHubBadge]: https://img.shields.io/badge/GitHub-181717?style=for-the-badge&logo=github&logoColor=white
[GitHubURL]: https://github.com/
[VSCodeBadge]: https://img.shields.io/badge/VSCode-0078D4?style=for-the-badge&logo=visualstudiocode&logoColor=white
[VSCodeURL]: https://code.visualstudio.com/
[SQLCBadge]: https://img.shields.io/badge/SQLC-336791?style=for-the-badge&logo=postgresql&logoColor=white
[SQLCURL]: https://sqlc.dev/
[GooseBadge]: https://img.shields.io/badge/Goose-FFD700?style=for-the-badge&logo=go&logoColor=black
[GooseURL]: https://pressly.github.io/goose/
[ActionsBadge]: https://img.shields.io/badge/GitHub%20Actions-2088FF?style=for-the-badge&logo=githubactions&logoColor=white
[ActionsURL]: https://github.com/features/actions

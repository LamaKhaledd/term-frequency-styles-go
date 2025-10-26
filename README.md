<a name="readme-top"></a>
<div align="center">
  <img src="third.PNG" alt="Project banner image">
  <br>
  <h1>Exercises in Programming Style –Go Implementations 💡</h1>
  <strong>Exploring multiple programming paradigms in Go — one problem, many perspectives 🚀</strong>
  <br><br>
  <img src="fourth.PNG" alt="Supporting image">
</div>

<br>

<div align="center">
  <p>
    Re-implementing the term-frequency task from Cristina Videira Lopes’ <em>Exercises in Programming Style</em> using Go.  
    This project demonstrates how constraints and idioms shape design choices across diverse programming styles.
  </p>
  <br>
  <a href="https://github.com/YourUserName/exercises-in-style-go"><strong>📘 View the Repository »</strong></a>
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
- Include tests and performance comparisons.

<br>



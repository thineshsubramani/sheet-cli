# Cheatsheet CLI Tool

**Cheatsheet CLI** is a fast, extensible command-line tool to fetch and display language or tech cheatsheets from simple YAML files. Perfect for DevOps, SREs, and developers who want quick access to syntax snippets without jumping into heavy docs.

## Features
- Supports multiple languages & tech stacks via YAML config
- Dynamic aggregation of sections with `all` keyword (combines all snippets)
- Alias support for quick language selection (e.g., `go`, `g`, `golang`)
- Fetch from remote URLs (GitHub raw) or local files seamlessly
- Minimal dependencies, pure Go implementation
- Detail flag to toggle syntax only or full description + syntax

## How To Use
```bash
sheet <language> <section> [-b backend-url] [-d]
````

* Example: `sheet go all` to get entire Go cheatsheet aggregated
* Use `-b` to override backend (defaults to GitHub raw URL)
* Use `-d` to show detailed descriptions alongside syntax

## Repo Structure

* `main.yaml` — Language registry with aliases and YAML paths
* `<language>/<language>.yaml` — Syntax snippets per language
* `sheet.go` — Main Go CLI program

## Why Use This

* Lightweight & fast lookup of syntax
* Easy to extend by adding new YAML files
* Great for SREs/DevOps wanting quick refreshers on languages or CLI commands

---

Built with ❤️ by Thinesh.
Contributions & issues welcome.

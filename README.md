# 🧠 sheet-cli — Your Terminal Cheat Sheet Sidekick

Simple Go CLI to fetch and display syntax snippets for popular programming languages like **Go**, **Python**, etc., from a remote YAML-based backend. Built-in alias support, in-memory caching, and colorized output. Fast, clean, and made for DevOps/SREs who live in terminals.

---

## 🌐 Backend YAML Source

All cheat sheets are hosted here:  
**https://raw.githubusercontent.com/thineshsubramani/cheatsheet/main/**

---

## ⚡ Cache System

Multi-layer caching for speed and offline support:
1. 🧠 **In-memory** – Fastest access after first load  
2. 💾 **Local file** – Cached YAML files stored locally  
3. 🌐 **Remote GitHub** – Fetched if not cached or forced refresh  

Use `-r` to refresh everything manually.

---

## 🚀 Usage

```bash
go run main.go <language> [section1,section2,...]
````

### ✅ Basic Examples

```bash
go run main.go go map,slice
go run main.go go help
go run main.go python all
```

### 🧽 Refresh All Cache

```bash
go run main.go --r go all
```

### 📖 Flags

| Flag  | Shortcut | Description                                   |
| ----- | -------- | --------------------------------------------- |
| `--b` | `-b`     | Custom backend base URL or local path         |
| `--d` | `-d`     | Show detailed output with description         |
| `--r` | `-r`     | Force refresh: reload everything from backend |

---

## 🧩 Aliases and Help

Every section can optionally define an alias (e.g. `map` → `m`) in the YAML.
Use `help` as section name to list all available sections with their aliases:

```bash
go run main.go go help
```

---

## 🎨 Output Colors

* 💙 Titles = Cyan
* 💛 Descriptions = Yellow
* 💚 Syntax/Code = Green

Built-in ANSI terminal coloring. Works great on:

* macOS/Linux terminals ✅
* Git Bash, WSL, MobaXterm on Windows ✅

---

## ✍️ Example Output

```bash
go run main.go go map
```

```
---- Map ----
Go maps basics

m := map[string]int{"a": 1, "b": 2}
fmt.Println(m["a"]) // access map
```

---

## 👊 That's It!

Just run, fetch, and hack. Update your YAMLs, and hit `-r` to sync changes instantly.
Made for SREs and CLI lovers who don’t wanna open a browser for every damn syntax question.

```

---

Let me know if you want to add:
- Support for custom local YAML paths
- Offline-only mode
- GitHub-style auto-update checker

Happy hacking, Thinesh 🤘
```

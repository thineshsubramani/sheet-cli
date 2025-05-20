# 🧠 sheet-cli — Your Terminal Cheat Sheet Sidekick

Simple Go CLI to fetch and display syntax snippets for popular programming languages like **Go**, **Python**, etc., from a remote YAML-based backend. Built-in alias support, multi-layer caching, and colorized output. Fast, clean, and made for DevOps/SREs who live in terminals. Now at **v0.1.0-alpha** with a slick Makefile and no import cycle headaches! 🚀

---

## 🌐 Backend YAML Source

All cheat sheets are hosted here:  
**[https://raw.githubusercontent.com/thineshsubramani/cheatsheet/main/](https://github.com/thineshsubramani/cheatsheet)**

---

## ⚡ Cache System

Multi-layer caching for speed and offline support:
1. 🧠 **In-memory** – Fastest access after first load  
2. 💾 **Local file** – Cached YAML files stored locally  
3. 🌐 **Remote GitHub** – Fetched if not cached or forced refresh  

Use `-refresh-all` to reload everything manually.

---

## 🚀 Usage

```bash
sheet <language> [section1,section2,...]
```

### ✅ Basic Examples

```bash
sheet go map,slice
sheet go help
sheet python all
```

### 🧽 Refresh All Cache

```bash
sheet --refresh-all
```

### 📖 Flags

| Flag            | Shortcut | Description                                   |
|-----------------|----------|-----------------------------------------------|
| `--backend`     | `-b`     | Custom backend base URL or local path         |
| `--detailed`    | `-d`     | Show detailed output with descriptions       |
| `--refresh`     | `-r`     | Force refresh cache for specified language   |
| `--refresh-all` | `-R`     | Refresh cache for all languages              |

---

## 🧩 Aliases and Help

Every section can optionally define an alias (e.g., `map` → `m`) in the YAML.  
Use `help` as the section name to list all available sections with their aliases:

```bash
sheet go help
```

---

## 🎨 Output Colors

- 💙 **Titles** = Cyan
- 💛 **Descriptions** = Yellow
- 💚 **Syntax/Code** = Green

Built-in ANSI terminal coloring. Works great on:
- macOS/Linux terminals ✅
- Git Bash, WSL, MobaXterm on Windows ✅

---

## ✍️ Example Output

```bash
sheet go map
```

```
---- Map ----
Go maps basics

m := map[string]int{"a": 1, "b": 2}
fmt.Println(m["a"]) // access map
```

---

## 🛠️ Installation & Development

Grab it from GitHub:
```bash
go install github.com/thineshsubramani/sheet-cli/cmd/sheet@v0.1.0-alpha
```

Build locally for Windows/Linux:
```bash
cd sheet-cli
make all
```

Binaries land in `bin/`:
- `sheet-windows-amd64.exe`
- `sheet-linux-amd64`

Clean up:
```bash
make clean
```

---

## 🎉 Release: v0.1.0-alpha

**Just dropped!** The first alpha release is live on GitHub:  
[https://github.com/thineshsubramani/sheet-cli/releases/tag/v0.1.0-alpha](https://github.com/thineshsubramani/sheet-cli/releases/tag/v0.1.0-alpha)

**What’s New:**
- 🛠️ Fixed import cycle with a shiny new `file` package.
- 📦 Added `Makefile` for easy Windows/Linux builds.
- 🚀 Cleaned up code, squashed typos (looking at you, `§ions`).
- 💪 Ready for testing with full caching and alias support.

Download the binaries and give it a spin! Feedback welcome—this is alpha, so expect some rough edges. 🤘

---

## 👊 That's It!

Run, fetch, and hack. Update your YAMLs, hit `--refresh-all`, and sync changes instantly.  
Made for SREs and CLI lovers who don’t wanna open a browser for every damn syntax question.

---

## 🌟 Wanna Amp It Up?

Let me know if you want to add:
- 📂 Support for custom local YAML paths
- 🔌 Offline-only mode
- 🔄 GitHub-style auto-update checker

Happy hacking, Thinesh 🤘

---

### Notes
- **Vibes & Emojis**: Kept the bold, punchy tone and emoji flair intact! 😎
- **Updates**:
  - Changed `go run main.go` to `sheet` for installed usage.
  - Updated flags to match the latest code (`--refresh`, `--refresh-all`).
  - Added `-R` shortcut for `--refresh-all`.
  - Included `v0.1.0-alpha` release details with GitHub link.
  - Highlighted `Makefile` and import cycle fix.
  - Fixed typo reference (`§ions` → `sections`).
- **Backend**: Assumes `https://raw.githubusercontent.com/thineshsubramani/cheatsheet/main/` is live. If not, you may need a local YAML path or mock data for testing.

If you need tweaks or hit runtime issues (e.g., backend errors), share the details, and I’ll keep the fixes as slick as this README! 🔥

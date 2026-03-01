
## 📌 AutoPortFinder (APF)

**AutoPortFinder** is a CLI tool to **find which processes are using which ports** — and optionally **kill them**.
Useful for developers, DevOps, or anyone who runs local servers and needs to free ports quickly.

---

## 🚀 User Guide

### ✅ Features

* 🔍 Scan all `LISTENING` TCP ports
* 🎯 Filter by port number or process name
* 📊 Display results in table, JSON, or plain format
* 🧨 Kill processes occupying ports (optional)

---

### 📦 Installation

```bash
go install github.com/your-username/autoportfinder@latest
```

Or clone and build manually:

```bash
git clone https://github.com/your-username/autoportfinder.git
cd autoportfinder
go build -o apf
```

---

### 🧪 Usage

```bash
./apf [flags]
```

#### 🔍 Scan open ports:

```bash
apf
```

#### 🎯 Filter by port or process name:

```bash
apf --port 8080
apf --process nginx
```

#### 🖨️ Output formats:

```bash
apf --format table      # Default
apf --format json
apf --format interface  # Simple plain-text
```

#### 🧨 Kill process on port:

```bash
apf --port 3000 --kill
```

> ⚠️ Requires root/admin privileges depending on your OS.

---

### 🔧 CLI Flags

| Flag            | Type   | Description                                 |
| --------------- | ------ | ------------------------------------------- |
| `-p, --port`    | int    | Filter by specific port (e.g. `8080`)       |
| `-r, --process` | string | Filter by process name                      |
| `-f, --format`  | string | Output format: `table`, `json`, `interface` |
| `--kill`        | bool   | Kill the found process(es)                  |

---

## ⚙️ Developer Guide

### 🧱 Architecture

The project follows a **Clean Architecture**-like separation of concerns:

```
autoportfinder/
├── cmd/                # CLI and flags (Cobra)
├── core/               # Runner (orchestrates Scanner, Output, Killer)
├── adapter/            # Concrete implementations of Scanner, Killer
├── output/             # Concrete Output formats
├── domain/             # Interfaces + shared DTOs (AppConfig, PortInfo)
└── main.go             # Entry point
```

---

### 🧩 Interfaces (domain/interfaces.go)

```go
type Scanner interface {
    Scan(cfg *AppConfig) ([]PortInfo, error)
}

type Output interface {
    Print([]PortInfo, format string) error
}

type Killer interface {
    Kill([]PortInfo) error
}
```

---

### 📁 Module Responsibilities

| Layer     | File                                  | Description                                          |
| --------- | ------------------------------------- | ---------------------------------------------------- |
| `cmd`     | `root.go`                             | CLI entrypoint (flags, command routing)              |
| `core`    | `runner.go`                           | Central orchestration logic                          |
| `adapter` | `scanner.go`                          | Uses gopsutil to collect open ports and process info |
|           | `killer.go`                           | Kills processes via PID                              |
| `output`  | `table.go`, `json.go`, `interface.go` | Different output renderers                           |
| `domain`  | `app_config.go`                       | Config structure                                     |
|           | `interfaces.go`                       | Domain interfaces                                    |

---

### 📚 Dependencies

* [`gopsutil`](https://github.com/shirou/gopsutil) — for fetching process/port info
* [`tablewriter`](https://github.com/olekukonko/tablewriter) — for pretty table output
* [`cobra`](https://github.com/spf13/cobra) — CLI framework

---

### 🧪 Run tests

> Unit tests can be added in each package.

```bash
go test ./...
```

---

### 🤝 Contributing

Pull requests and issues are welcome! Please lint and format your code with `go fmt`.

---

### 📜 License

MIT License

---

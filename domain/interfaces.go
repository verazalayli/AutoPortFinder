package domain

// PortInfo describes a single open port with metadata
type PortInfo struct {
	Port     int
	PID      int
	ProcName string
}

// Output defines a generic output printer (table, json, interface view, etc.)
type Output interface {
	Print(results []PortInfo, format string) error
}

// Scanner is the abstraction for scanning ports/processes
type Scanner interface {
	Scan(cfg *AppConfig) ([]PortInfo, error)
}

// Killer defines an interface for terminating a process by PID
type Killer interface {
	KillProcess(info PortInfo) error
}

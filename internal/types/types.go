package types

type PortInfo struct {
	Port     int
	PID      int
	ProcName string
}

type AppConfig struct {
	PortFilter    int
	ProcessFilter string
	OutputFormat  string
	KillFlag      bool
	Watch         bool
}

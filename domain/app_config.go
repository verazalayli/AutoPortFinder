package domain

type AppConfig struct {
	PortFilter    int    // e.g. 8080 — to filter by port number
	ProcessFilter string // e.g. "nginx" — to filter by process name
	OutputFormat  string // e.g. "json", "table", "interface"
	KillFlag      bool   // if true, try to kill matching process(es)
}

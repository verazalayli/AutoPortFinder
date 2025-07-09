package config

type AppConfig struct {
	PortFilter    int
	ProcessFilter string
	OutputFormat  string
	KillFlag      bool
	Watch         bool
}

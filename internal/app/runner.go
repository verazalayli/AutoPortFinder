package app

import (
	"fmt"
	"github.com/verazalyali/internal/output"
	"github.com/verazalyali/internal/ports"
	"github.com/verazalyali/internal/types"
)

func Run(cfg *types.AppConfig) error {
	results, err := ports.ScanPorts(cfg)
	if err != nil {
		return fmt.Errorf("scan failed: %v", err)
	}

	if len(results) == 0 {
		fmt.Println("No matching ports found.")
		return nil
	}

	switch cfg.OutputFormat {
	case "json":
		return output.OutputJSON(results)
	case "table":
		output.OutputTable(results)
	case "interface":
		output.OutputInterface(results)
	default:
		return fmt.Errorf("unknown output format: %s", cfg.OutputFormat)
	}

	if cfg.KillFlag {
		// Удаление процессов — не реализовано
	}

	return nil
}

package app

import (
	"fmt"
	killer "github.com/verazalyali/internal/killer"
	"github.com/verazalyali/internal/ports"
	"github.com/verazalyali/internal/types"
)

func Run(cfg *types.AppConfig) error {
	results, err := ports.ScanPorts(cfg)
	if err != nil {
		return fmt.Errorf("scan error: %w", err)
	}

	if cfg.KillFlag {
		if cfg.PortFilter == 0 {
			return fmt.Errorf("to use --kill you must specify a port with --port")
		}
		return killer.KillByPort(results, cfg.PortFilter)
	}

	return nil
}

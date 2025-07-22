package adapter

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/verazalyali/domain"
)

type PortScanner struct{}

func NewScanner() domain.Scanner {
	return &PortScanner{}
}

func (s *PortScanner) Scan(cfg *domain.AppConfig) ([]domain.PortInfo, error) {
	connections, err := net.Connections("tcp")
	if err != nil {
		return nil, fmt.Errorf("failed to get TCP connections: %w", err)
	}

	var results []domain.PortInfo

	for _, conn := range connections {
		if conn.Status != "LISTEN" {
			continue
		}

		// filter by port
		if cfg.PortFilter > 0 && int(conn.Laddr.Port) != cfg.PortFilter {
			continue
		}

		proc, err := process.NewProcess(conn.Pid)
		if err != nil {
			fmt.Printf("Can't get process for pid %d: %v\n", conn.Pid, err)
			continue
		}

		name, err := proc.Name()
		if err != nil {
			fmt.Printf("Can't get name for pid %d: %v\n", conn.Pid, err)
			continue
		}

		// filter by process name
		if cfg.ProcessFilter != "" && name != cfg.ProcessFilter {
			continue
		}

		results = append(results, domain.PortInfo{
			Port:     int(conn.Laddr.Port),
			PID:      int(conn.Pid),
			ProcName: name,
		})
	}

	return results, nil
}

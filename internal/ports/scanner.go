package ports

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/verazalyali/internal/config"
	"os"
)

// PortInfo describes information about connection
type PortInfo struct {
	Port     uint32
	PID      int32
	ProcName string
}

func outputJSON(results []PortInfo) error {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	fmt.Println(string(data))
	return nil
}

func outputTable(results []PortInfo) {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Port", "PID", "Process"})

	for _, r := range results {
		table.Append([]string{
			fmt.Sprintf("%d", r.Port),
			fmt.Sprintf("%d", r.PID),
			r.ProcName,
		})
	}

	table.Render()
}

func outputInterface(results []PortInfo) {
	for _, r := range results {
		fmt.Printf("Port: %d\nPID: %d\nProcess: %s\n---\n", r.Port, r.PID, r.ProcName)
	}
}

// RunScan search for all listening ports and processes
func RunScan(cfg *config.AppConfig) error {
	connections, err := net.Connections("tcp")
	if err != nil {
		return err
	}

	var results []PortInfo

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
			fmt.Printf("warning: can't get process for pid %d\n", conn.Pid)
			continue
		}

		name, err := proc.Name()
		if err != nil {
			fmt.Printf("warning: can't get process name for pid %d\n", conn.Pid)
			continue
		}

		// filter by name
		if cfg.ProcessFilter != "" && name != cfg.ProcessFilter {
			continue
		}

		results = append(results, PortInfo{
			Port:     conn.Laddr.Port,
			PID:      conn.Pid,
			ProcName: name,
		})
	}

	switch cfg.OutputFormat {
	case "json":
		return outputJSON(results)
	case "table":
		outputTable(results)
	case "interface":
		outputInterface(results)
	default:
		return fmt.Errorf("unknown output format: %s", cfg.OutputFormat)
	}

	return nil
}

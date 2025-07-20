package killerpackage

import (
	"fmt"
	"github.com/verazalyali/internal/types"
	"os"
	"syscall"
)

func KillByPort(results []types.PortInfo, port int) error {
	var killed bool

	for _, r := range results {
		if r.Port == port {
			process, err := os.FindProcess(r.PID)
			if err != nil {
				fmt.Printf("⚠️  Can't find process PID %d: %v\n", r.PID, err)
				continue
			}

			fmt.Printf("❌ Killing process %s (PID %d) on port %d...\n", r.ProcName, r.PID, r.Port)
			err = process.Signal(syscall.SIGKILL)
			if err != nil {
				fmt.Printf("⚠️  Failed to kill PID %d: %v\n", r.PID, err)
			} else {
				fmt.Printf("✅ Process %d killed.\n", r.PID)
				killed = true
			}
		}
	}

	if !killed {
		return fmt.Errorf("no process found listening on port %d", port)
	}
	return nil
}

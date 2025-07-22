package adapter

import (
	"fmt"
	"github.com/verazalyali/domain"
	"syscall"
)

type killerAdapter struct{}

func NewKiller() domain.Killer {
	return &killerAdapter{}
}

// KillProcess завершает процесс по PID
func (k *killerAdapter) KillProcess(info domain.PortInfo) error {
	err := syscall.Kill(info.PID, syscall.SIGKILL)
	if err != nil {
		return fmt.Errorf("failed to kill process %d: %w", info.PID, err)
	}
	fmt.Printf("✅ Process %d (%s) on port %d killed\n", info.PID, info.ProcName, info.Port)
	return nil
}

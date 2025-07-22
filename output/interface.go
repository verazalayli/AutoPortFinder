package output

import (
	"fmt"
	"github.com/verazalyali/domain"
)

type InterfaceOutput struct{}

func NewInterfaceOutput() *InterfaceOutput {
	return &InterfaceOutput{}
}

func (i *InterfaceOutput) Print(results []domain.PortInfo, _ string) error {
	for _, r := range results {
		fmt.Printf("Port: %d\nPID: %d\nProcess: %s\n---\n", r.Port, r.PID, r.ProcName)
	}
	return nil
}

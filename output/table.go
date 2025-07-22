package output

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/verazalyali/domain"
	"os"
)

type TableOutput struct{}

func NewTableOutput() *TableOutput {
	return &TableOutput{}
}

func (t *TableOutput) Print(results []domain.PortInfo, _ string) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Port", "PID", "Process"})

	for _, r := range results {
		err := table.Append([]interface{}{
			r.Port,
			r.PID,
			r.ProcName,
		})
		if err != nil {
			return fmt.Errorf("failed to append row: %w", err)
		}
	}

	if err := table.Render(); err != nil {
		return fmt.Errorf("failed to render table: %w", err)
	}

	return nil

}

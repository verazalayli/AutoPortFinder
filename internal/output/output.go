package output

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/verazalyali/internal/types"
	"os"
)

func OutputJSON(results []types.PortInfo) error {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	fmt.Println(string(data))
	return nil
}

func OutputTable(results []types.PortInfo) {
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

func OutputInterface(results []types.PortInfo) {
	for _, r := range results {
		fmt.Printf("Port: %d\nPID: %d\nProcess: %s\n---\n", r.Port, r.PID, r.ProcName)
	}
}

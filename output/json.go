package output

import (
	"encoding/json"
	"fmt"
	"github.com/verazalyali/domain"
)

type JSONOutput struct{}

func NewJSONOutput() *JSONOutput {
	return &JSONOutput{}
}

func (j *JSONOutput) Print(results []domain.PortInfo, _ string) error {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	fmt.Println(string(data))
	return nil
}

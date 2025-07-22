package core

import (
	"fmt"
	"github.com/verazalyali/domain"
)

type Runner struct {
	Scanner domain.Scanner
	Output  domain.Output
	Killer  domain.Killer
}

func NewRunner(scanner domain.Scanner, output domain.Output, killer domain.Killer) *Runner {
	return &Runner{
		Scanner: scanner,
		Output:  output,
		Killer:  killer,
	}
}

func (r *Runner) Run(cfg *domain.AppConfig) error {
	results, err := r.Scanner.Scan(cfg)
	if err != nil {
		return err
	}

	if err := r.Output.Print(results, cfg.OutputFormat); err != nil {
		return err
	}

	if cfg.KillFlag {
		for _, info := range results {
			if err := r.Killer.KillProcess(info); err != nil {
				fmt.Printf("⚠️ Failed to kill process %d on port %d: %v\n", info.PID, info.Port, err)
			}
		}
	}

	return nil
}

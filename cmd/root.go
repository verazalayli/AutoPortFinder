package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/verazalyali/internal/ports"
	"github.com/verazalyali/internal/types"
	"os"
)

var (
	cfg     = &types.AppConfig{}
	rootCmd = &cobra.Command{
		Use:   "autoportfinder",
		Short: "Show which processes are using which port",
		RunE: func(cmd *cobra.Command, args []string) error {
			return ports.RunScan(cfg)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().IntVarP(&cfg.PortFilter, "port", "p", 0, "Filter by port (8080 or other)")
	rootCmd.PersistentFlags().StringVarP(&cfg.ProcessFilter, "process", "r", "", "Filter by process name (nginx or other)")
	rootCmd.PersistentFlags().StringVarP(&cfg.OutputFormat, "format", "f", "table", "Output format: table, json, csv")
	rootCmd.PersistentFlags().BoolVar(&cfg.KillFlag, "kill", false, "Kill process on port")

	// here you can add commands
	// rootCmd.AddCommand(...)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

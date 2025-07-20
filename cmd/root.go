package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	runner "github.com/verazalyali/internal/app"
	"github.com/verazalyali/internal/types"
	"os"
)

var (
	cfg     = &types.AppConfig{}
	rootCmd = &cobra.Command{
		Use:   "apf", // или "autoportfinder"
		Short: "Show which processes are using which ports",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runner.Run(cfg)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().IntVarP(&cfg.PortFilter, "port", "p", 0, "Filter by port (e.g., 8080)")
	rootCmd.PersistentFlags().StringVarP(&cfg.ProcessFilter, "process", "r", "", "Filter by process name (e.g., nginx)")
	rootCmd.PersistentFlags().StringVarP(&cfg.OutputFormat, "format", "f", "table", "Output format: table, json, interface")
	rootCmd.PersistentFlags().BoolVar(&cfg.KillFlag, "kill", false, "Kill process on the specified port")

	// here you can add commands
	// rootCmd.AddCommand(...)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

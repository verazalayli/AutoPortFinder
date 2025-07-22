package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/verazalyali/adapter"
	"github.com/verazalyali/core"
	"github.com/verazalyali/domain"
	"github.com/verazalyali/output"
	"os"
)

var (
	cfg = &domain.AppConfig{}
)

var rootCmd = &cobra.Command{
	Use:   "apf",
	Short: "AutoPortFinder: find and optionally kill processes using ports",
	RunE: func(cmd *cobra.Command, args []string) error {
		scanner := adapter.NewScanner()
		killer := adapter.NewKiller()
		var out domain.Output

		switch cfg.OutputFormat {
		case "json":
			out = output.NewJSONOutput()
		case "table":
			out = output.NewTableOutput()
		case "interface":
			out = output.NewInterfaceOutput()
		default:
			return fmt.Errorf("unknown output format: %s", cfg.OutputFormat)
		}

		runner := core.NewRunner(scanner, out, killer)
		return runner.Run(cfg)
	},
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&cfg.PortFilter, "port", "p", 0, "Filter by port (e.g., 8080)")
	rootCmd.PersistentFlags().StringVarP(&cfg.ProcessFilter, "process", "r", "", "Filter by process name (e.g., nginx)")
	rootCmd.PersistentFlags().StringVarP(&cfg.OutputFormat, "format", "f", "table", "Output format: table, json, interface")
	rootCmd.PersistentFlags().BoolVar(&cfg.KillFlag, "kill", false, "Kill process on the specified port")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "❌", err)
		os.Exit(1)
	}
}

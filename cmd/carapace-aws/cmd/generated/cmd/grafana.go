package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafanaCmd = &cobra.Command{
	Use:   "grafana",
	Short: "Amazon Managed Grafana is a fully managed and secure data visualization service that you can use to instantly query, correlate, and visualize operational metrics, logs, and traces from multiple sources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafanaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(grafanaCmd).Standalone()

	})
	rootCmd.AddCommand(grafanaCmd)
}

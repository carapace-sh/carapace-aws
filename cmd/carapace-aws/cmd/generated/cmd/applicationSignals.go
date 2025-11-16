package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignalsCmd = &cobra.Command{
	Use:   "application-signals",
	Short: "Use CloudWatch Application Signals for comprehensive observability of your cloud-based applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignalsCmd).Standalone()

	rootCmd.AddCommand(applicationSignalsCmd)
}

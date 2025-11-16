package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsightsCmd = &cobra.Command{
	Use:   "application-insights",
	Short: "Amazon CloudWatch Application Insights",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsightsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationInsightsCmd).Standalone()

	})
	rootCmd.AddCommand(applicationInsightsCmd)
}

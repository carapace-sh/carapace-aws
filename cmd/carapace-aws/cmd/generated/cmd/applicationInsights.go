package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsightsCmd = &cobra.Command{
	Use:   "application-insights",
	Short: "Amazon CloudWatch Application Insights\n\nAmazon CloudWatch Application Insights is a service that helps you detect common problems with your applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsightsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationInsightsCmd).Standalone()

	})
	rootCmd.AddCommand(applicationInsightsCmd)
}

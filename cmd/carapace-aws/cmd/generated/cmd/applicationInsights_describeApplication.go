package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_describeApplicationCmd = &cobra.Command{
	Use:   "describe-application",
	Short: "Describes the application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_describeApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationInsights_describeApplicationCmd).Standalone()

		applicationInsights_describeApplicationCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the resource group owner.")
		applicationInsights_describeApplicationCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
		applicationInsights_describeApplicationCmd.MarkFlagRequired("resource-group-name")
	})
	applicationInsightsCmd.AddCommand(applicationInsights_describeApplicationCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_deleteApplicationCmd = &cobra.Command{
	Use:   "delete-application",
	Short: "Removes the specified application from monitoring.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_deleteApplicationCmd).Standalone()

	applicationInsights_deleteApplicationCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
	applicationInsights_deleteApplicationCmd.MarkFlagRequired("resource-group-name")
	applicationInsightsCmd.AddCommand(applicationInsights_deleteApplicationCmd)
}

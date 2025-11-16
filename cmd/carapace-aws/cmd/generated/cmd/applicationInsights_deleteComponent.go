package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_deleteComponentCmd = &cobra.Command{
	Use:   "delete-component",
	Short: "Ungroups a custom component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_deleteComponentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationInsights_deleteComponentCmd).Standalone()

		applicationInsights_deleteComponentCmd.Flags().String("component-name", "", "The name of the component.")
		applicationInsights_deleteComponentCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
		applicationInsights_deleteComponentCmd.MarkFlagRequired("component-name")
		applicationInsights_deleteComponentCmd.MarkFlagRequired("resource-group-name")
	})
	applicationInsightsCmd.AddCommand(applicationInsights_deleteComponentCmd)
}

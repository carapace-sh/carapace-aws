package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_updateComponentCmd = &cobra.Command{
	Use:   "update-component",
	Short: "Updates the custom component name and/or the list of resources that make up the component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_updateComponentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationInsights_updateComponentCmd).Standalone()

		applicationInsights_updateComponentCmd.Flags().String("component-name", "", "The name of the component.")
		applicationInsights_updateComponentCmd.Flags().String("new-component-name", "", "The new name of the component.")
		applicationInsights_updateComponentCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
		applicationInsights_updateComponentCmd.Flags().String("resource-list", "", "The list of resource ARNs that belong to the component.")
		applicationInsights_updateComponentCmd.MarkFlagRequired("component-name")
		applicationInsights_updateComponentCmd.MarkFlagRequired("resource-group-name")
	})
	applicationInsightsCmd.AddCommand(applicationInsights_updateComponentCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_createComponentCmd = &cobra.Command{
	Use:   "create-component",
	Short: "Creates a custom component by grouping similar standalone instances to monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_createComponentCmd).Standalone()

	applicationInsights_createComponentCmd.Flags().String("component-name", "", "The name of the component.")
	applicationInsights_createComponentCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
	applicationInsights_createComponentCmd.Flags().String("resource-list", "", "The list of resource ARNs that belong to the component.")
	applicationInsights_createComponentCmd.MarkFlagRequired("component-name")
	applicationInsights_createComponentCmd.MarkFlagRequired("resource-group-name")
	applicationInsights_createComponentCmd.MarkFlagRequired("resource-list")
	applicationInsightsCmd.AddCommand(applicationInsights_createComponentCmd)
}

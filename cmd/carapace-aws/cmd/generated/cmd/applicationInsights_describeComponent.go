package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_describeComponentCmd = &cobra.Command{
	Use:   "describe-component",
	Short: "Describes a component and lists the resources that are grouped together in a component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_describeComponentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationInsights_describeComponentCmd).Standalone()

		applicationInsights_describeComponentCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the resource group owner.")
		applicationInsights_describeComponentCmd.Flags().String("component-name", "", "The name of the component.")
		applicationInsights_describeComponentCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
		applicationInsights_describeComponentCmd.MarkFlagRequired("component-name")
		applicationInsights_describeComponentCmd.MarkFlagRequired("resource-group-name")
	})
	applicationInsightsCmd.AddCommand(applicationInsights_describeComponentCmd)
}

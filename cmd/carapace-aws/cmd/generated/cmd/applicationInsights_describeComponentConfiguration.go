package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_describeComponentConfigurationCmd = &cobra.Command{
	Use:   "describe-component-configuration",
	Short: "Describes the monitoring configuration of the component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_describeComponentConfigurationCmd).Standalone()

	applicationInsights_describeComponentConfigurationCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the resource group owner.")
	applicationInsights_describeComponentConfigurationCmd.Flags().String("component-name", "", "The name of the component.")
	applicationInsights_describeComponentConfigurationCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
	applicationInsights_describeComponentConfigurationCmd.MarkFlagRequired("component-name")
	applicationInsights_describeComponentConfigurationCmd.MarkFlagRequired("resource-group-name")
	applicationInsightsCmd.AddCommand(applicationInsights_describeComponentConfigurationCmd)
}

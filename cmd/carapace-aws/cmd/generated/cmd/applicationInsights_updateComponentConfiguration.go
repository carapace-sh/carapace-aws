package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_updateComponentConfigurationCmd = &cobra.Command{
	Use:   "update-component-configuration",
	Short: "Updates the monitoring configurations for the component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_updateComponentConfigurationCmd).Standalone()

	applicationInsights_updateComponentConfigurationCmd.Flags().String("auto-config-enabled", "", "Automatically configures the component by applying the recommended configurations.")
	applicationInsights_updateComponentConfigurationCmd.Flags().String("component-configuration", "", "The configuration settings of the component.")
	applicationInsights_updateComponentConfigurationCmd.Flags().String("component-name", "", "The name of the component.")
	applicationInsights_updateComponentConfigurationCmd.Flags().String("monitor", "", "Indicates whether the application component is monitored.")
	applicationInsights_updateComponentConfigurationCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
	applicationInsights_updateComponentConfigurationCmd.Flags().String("tier", "", "The tier of the application component.")
	applicationInsights_updateComponentConfigurationCmd.MarkFlagRequired("component-name")
	applicationInsights_updateComponentConfigurationCmd.MarkFlagRequired("resource-group-name")
	applicationInsightsCmd.AddCommand(applicationInsights_updateComponentConfigurationCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_updateWorkloadCmd = &cobra.Command{
	Use:   "update-workload",
	Short: "Adds a workload to a component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_updateWorkloadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationInsights_updateWorkloadCmd).Standalone()

		applicationInsights_updateWorkloadCmd.Flags().String("component-name", "", "The name of the component.")
		applicationInsights_updateWorkloadCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
		applicationInsights_updateWorkloadCmd.Flags().String("workload-configuration", "", "The configuration settings of the workload.")
		applicationInsights_updateWorkloadCmd.Flags().String("workload-id", "", "The ID of the workload.")
		applicationInsights_updateWorkloadCmd.MarkFlagRequired("component-name")
		applicationInsights_updateWorkloadCmd.MarkFlagRequired("resource-group-name")
		applicationInsights_updateWorkloadCmd.MarkFlagRequired("workload-configuration")
	})
	applicationInsightsCmd.AddCommand(applicationInsights_updateWorkloadCmd)
}

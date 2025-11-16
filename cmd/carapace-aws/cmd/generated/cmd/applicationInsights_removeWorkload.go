package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_removeWorkloadCmd = &cobra.Command{
	Use:   "remove-workload",
	Short: "Remove workload from a component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_removeWorkloadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationInsights_removeWorkloadCmd).Standalone()

		applicationInsights_removeWorkloadCmd.Flags().String("component-name", "", "The name of the component.")
		applicationInsights_removeWorkloadCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
		applicationInsights_removeWorkloadCmd.Flags().String("workload-id", "", "The ID of the workload.")
		applicationInsights_removeWorkloadCmd.MarkFlagRequired("component-name")
		applicationInsights_removeWorkloadCmd.MarkFlagRequired("resource-group-name")
		applicationInsights_removeWorkloadCmd.MarkFlagRequired("workload-id")
	})
	applicationInsightsCmd.AddCommand(applicationInsights_removeWorkloadCmd)
}

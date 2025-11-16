package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_addWorkloadCmd = &cobra.Command{
	Use:   "add-workload",
	Short: "Adds a workload to a component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_addWorkloadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationInsights_addWorkloadCmd).Standalone()

		applicationInsights_addWorkloadCmd.Flags().String("component-name", "", "The name of the component.")
		applicationInsights_addWorkloadCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
		applicationInsights_addWorkloadCmd.Flags().String("workload-configuration", "", "The configuration settings of the workload.")
		applicationInsights_addWorkloadCmd.MarkFlagRequired("component-name")
		applicationInsights_addWorkloadCmd.MarkFlagRequired("resource-group-name")
		applicationInsights_addWorkloadCmd.MarkFlagRequired("workload-configuration")
	})
	applicationInsightsCmd.AddCommand(applicationInsights_addWorkloadCmd)
}

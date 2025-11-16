package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_describeWorkloadCmd = &cobra.Command{
	Use:   "describe-workload",
	Short: "Describes a workload and its configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_describeWorkloadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationInsights_describeWorkloadCmd).Standalone()

		applicationInsights_describeWorkloadCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the workload owner.")
		applicationInsights_describeWorkloadCmd.Flags().String("component-name", "", "The name of the component.")
		applicationInsights_describeWorkloadCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
		applicationInsights_describeWorkloadCmd.Flags().String("workload-id", "", "The ID of the workload.")
		applicationInsights_describeWorkloadCmd.MarkFlagRequired("component-name")
		applicationInsights_describeWorkloadCmd.MarkFlagRequired("resource-group-name")
		applicationInsights_describeWorkloadCmd.MarkFlagRequired("workload-id")
	})
	applicationInsightsCmd.AddCommand(applicationInsights_describeWorkloadCmd)
}

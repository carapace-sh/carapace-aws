package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_listWorkloadsCmd = &cobra.Command{
	Use:   "list-workloads",
	Short: "Lists the workloads that are configured on a given component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_listWorkloadsCmd).Standalone()

	applicationInsights_listWorkloadsCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the owner of the workload.")
	applicationInsights_listWorkloadsCmd.Flags().String("component-name", "", "The name of the component.")
	applicationInsights_listWorkloadsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	applicationInsights_listWorkloadsCmd.Flags().String("next-token", "", "The token to request the next page of results.")
	applicationInsights_listWorkloadsCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
	applicationInsights_listWorkloadsCmd.MarkFlagRequired("component-name")
	applicationInsights_listWorkloadsCmd.MarkFlagRequired("resource-group-name")
	applicationInsightsCmd.AddCommand(applicationInsights_listWorkloadsCmd)
}

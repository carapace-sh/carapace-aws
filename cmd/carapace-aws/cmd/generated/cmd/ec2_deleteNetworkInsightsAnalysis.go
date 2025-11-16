package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteNetworkInsightsAnalysisCmd = &cobra.Command{
	Use:   "delete-network-insights-analysis",
	Short: "Deletes the specified network insights analysis.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteNetworkInsightsAnalysisCmd).Standalone()

	ec2_deleteNetworkInsightsAnalysisCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteNetworkInsightsAnalysisCmd.Flags().String("network-insights-analysis-id", "", "The ID of the network insights analysis.")
	ec2_deleteNetworkInsightsAnalysisCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteNetworkInsightsAnalysisCmd.MarkFlagRequired("network-insights-analysis-id")
	ec2_deleteNetworkInsightsAnalysisCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_deleteNetworkInsightsAnalysisCmd)
}

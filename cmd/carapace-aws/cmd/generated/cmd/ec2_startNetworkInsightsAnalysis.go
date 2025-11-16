package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_startNetworkInsightsAnalysisCmd = &cobra.Command{
	Use:   "start-network-insights-analysis",
	Short: "Starts analyzing the specified path.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_startNetworkInsightsAnalysisCmd).Standalone()

	ec2_startNetworkInsightsAnalysisCmd.Flags().String("additional-accounts", "", "The member accounts that contain resources that the path can traverse.")
	ec2_startNetworkInsightsAnalysisCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_startNetworkInsightsAnalysisCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_startNetworkInsightsAnalysisCmd.Flags().String("filter-in-arns", "", "The Amazon Resource Names (ARN) of the resources that the path must traverse.")
	ec2_startNetworkInsightsAnalysisCmd.Flags().String("filter-out-arns", "", "The Amazon Resource Names (ARN) of the resources that the path will ignore.")
	ec2_startNetworkInsightsAnalysisCmd.Flags().String("network-insights-path-id", "", "The ID of the path.")
	ec2_startNetworkInsightsAnalysisCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_startNetworkInsightsAnalysisCmd.Flags().String("tag-specifications", "", "The tags to apply.")
	ec2_startNetworkInsightsAnalysisCmd.MarkFlagRequired("client-token")
	ec2_startNetworkInsightsAnalysisCmd.MarkFlagRequired("network-insights-path-id")
	ec2_startNetworkInsightsAnalysisCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_startNetworkInsightsAnalysisCmd)
}

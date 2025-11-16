package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_startNetworkInsightsAccessScopeAnalysisCmd = &cobra.Command{
	Use:   "start-network-insights-access-scope-analysis",
	Short: "Starts analyzing the specified Network Access Scope.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_startNetworkInsightsAccessScopeAnalysisCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_startNetworkInsightsAccessScopeAnalysisCmd).Standalone()

		ec2_startNetworkInsightsAccessScopeAnalysisCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_startNetworkInsightsAccessScopeAnalysisCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_startNetworkInsightsAccessScopeAnalysisCmd.Flags().String("network-insights-access-scope-id", "", "The ID of the Network Access Scope.")
		ec2_startNetworkInsightsAccessScopeAnalysisCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_startNetworkInsightsAccessScopeAnalysisCmd.Flags().String("tag-specifications", "", "The tags to apply.")
		ec2_startNetworkInsightsAccessScopeAnalysisCmd.MarkFlagRequired("client-token")
		ec2_startNetworkInsightsAccessScopeAnalysisCmd.MarkFlagRequired("network-insights-access-scope-id")
		ec2_startNetworkInsightsAccessScopeAnalysisCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_startNetworkInsightsAccessScopeAnalysisCmd)
}

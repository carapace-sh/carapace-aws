package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeNetworkInsightsAccessScopeAnalysesCmd = &cobra.Command{
	Use:   "describe-network-insights-access-scope-analyses",
	Short: "Describes the specified Network Access Scope analyses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeNetworkInsightsAccessScopeAnalysesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeNetworkInsightsAccessScopeAnalysesCmd).Standalone()

		ec2_describeNetworkInsightsAccessScopeAnalysesCmd.Flags().String("analysis-start-time-begin", "", "Filters the results based on the start time.")
		ec2_describeNetworkInsightsAccessScopeAnalysesCmd.Flags().String("analysis-start-time-end", "", "Filters the results based on the start time.")
		ec2_describeNetworkInsightsAccessScopeAnalysesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeNetworkInsightsAccessScopeAnalysesCmd.Flags().String("filters", "", "There are no supported filters.")
		ec2_describeNetworkInsightsAccessScopeAnalysesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeNetworkInsightsAccessScopeAnalysesCmd.Flags().String("network-insights-access-scope-analysis-ids", "", "The IDs of the Network Access Scope analyses.")
		ec2_describeNetworkInsightsAccessScopeAnalysesCmd.Flags().String("network-insights-access-scope-id", "", "The ID of the Network Access Scope.")
		ec2_describeNetworkInsightsAccessScopeAnalysesCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeNetworkInsightsAccessScopeAnalysesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeNetworkInsightsAccessScopeAnalysesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeNetworkInsightsAccessScopeAnalysesCmd)
}

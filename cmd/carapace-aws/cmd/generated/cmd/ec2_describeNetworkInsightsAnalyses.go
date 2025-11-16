package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeNetworkInsightsAnalysesCmd = &cobra.Command{
	Use:   "describe-network-insights-analyses",
	Short: "Describes one or more of your network insights analyses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeNetworkInsightsAnalysesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeNetworkInsightsAnalysesCmd).Standalone()

		ec2_describeNetworkInsightsAnalysesCmd.Flags().String("analysis-end-time", "", "The time when the network insights analyses ended.")
		ec2_describeNetworkInsightsAnalysesCmd.Flags().String("analysis-start-time", "", "The time when the network insights analyses started.")
		ec2_describeNetworkInsightsAnalysesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeNetworkInsightsAnalysesCmd.Flags().String("filters", "", "The filters.")
		ec2_describeNetworkInsightsAnalysesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeNetworkInsightsAnalysesCmd.Flags().String("network-insights-analysis-ids", "", "The ID of the network insights analyses.")
		ec2_describeNetworkInsightsAnalysesCmd.Flags().String("network-insights-path-id", "", "The ID of the path.")
		ec2_describeNetworkInsightsAnalysesCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeNetworkInsightsAnalysesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeNetworkInsightsAnalysesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeNetworkInsightsAnalysesCmd)
}

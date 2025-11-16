package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeNetworkInsightsPathsCmd = &cobra.Command{
	Use:   "describe-network-insights-paths",
	Short: "Describes one or more of your paths.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeNetworkInsightsPathsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeNetworkInsightsPathsCmd).Standalone()

		ec2_describeNetworkInsightsPathsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeNetworkInsightsPathsCmd.Flags().String("filters", "", "The filters.")
		ec2_describeNetworkInsightsPathsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeNetworkInsightsPathsCmd.Flags().String("network-insights-path-ids", "", "The IDs of the paths.")
		ec2_describeNetworkInsightsPathsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeNetworkInsightsPathsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeNetworkInsightsPathsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeNetworkInsightsPathsCmd)
}

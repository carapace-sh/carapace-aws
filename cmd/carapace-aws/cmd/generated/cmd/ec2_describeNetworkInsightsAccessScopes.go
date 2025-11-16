package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeNetworkInsightsAccessScopesCmd = &cobra.Command{
	Use:   "describe-network-insights-access-scopes",
	Short: "Describes the specified Network Access Scopes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeNetworkInsightsAccessScopesCmd).Standalone()

	ec2_describeNetworkInsightsAccessScopesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeNetworkInsightsAccessScopesCmd.Flags().String("filters", "", "There are no supported filters.")
	ec2_describeNetworkInsightsAccessScopesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_describeNetworkInsightsAccessScopesCmd.Flags().String("network-insights-access-scope-ids", "", "The IDs of the Network Access Scopes.")
	ec2_describeNetworkInsightsAccessScopesCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_describeNetworkInsightsAccessScopesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeNetworkInsightsAccessScopesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeNetworkInsightsAccessScopesCmd)
}

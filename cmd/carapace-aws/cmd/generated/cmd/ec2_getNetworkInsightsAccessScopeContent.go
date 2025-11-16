package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getNetworkInsightsAccessScopeContentCmd = &cobra.Command{
	Use:   "get-network-insights-access-scope-content",
	Short: "Gets the content for the specified Network Access Scope.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getNetworkInsightsAccessScopeContentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getNetworkInsightsAccessScopeContentCmd).Standalone()

		ec2_getNetworkInsightsAccessScopeContentCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getNetworkInsightsAccessScopeContentCmd.Flags().String("network-insights-access-scope-id", "", "The ID of the Network Access Scope.")
		ec2_getNetworkInsightsAccessScopeContentCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getNetworkInsightsAccessScopeContentCmd.MarkFlagRequired("network-insights-access-scope-id")
		ec2_getNetworkInsightsAccessScopeContentCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_getNetworkInsightsAccessScopeContentCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createNetworkInsightsAccessScopeCmd = &cobra.Command{
	Use:   "create-network-insights-access-scope",
	Short: "Creates a Network Access Scope.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createNetworkInsightsAccessScopeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createNetworkInsightsAccessScopeCmd).Standalone()

		ec2_createNetworkInsightsAccessScopeCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_createNetworkInsightsAccessScopeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createNetworkInsightsAccessScopeCmd.Flags().String("exclude-paths", "", "The paths to exclude.")
		ec2_createNetworkInsightsAccessScopeCmd.Flags().String("match-paths", "", "The paths to match.")
		ec2_createNetworkInsightsAccessScopeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createNetworkInsightsAccessScopeCmd.Flags().String("tag-specifications", "", "The tags to apply.")
		ec2_createNetworkInsightsAccessScopeCmd.MarkFlagRequired("client-token")
		ec2_createNetworkInsightsAccessScopeCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_createNetworkInsightsAccessScopeCmd)
}

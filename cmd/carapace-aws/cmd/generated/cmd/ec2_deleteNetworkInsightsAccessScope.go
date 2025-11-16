package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteNetworkInsightsAccessScopeCmd = &cobra.Command{
	Use:   "delete-network-insights-access-scope",
	Short: "Deletes the specified Network Access Scope.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteNetworkInsightsAccessScopeCmd).Standalone()

	ec2_deleteNetworkInsightsAccessScopeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteNetworkInsightsAccessScopeCmd.Flags().String("network-insights-access-scope-id", "", "The ID of the Network Access Scope.")
	ec2_deleteNetworkInsightsAccessScopeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteNetworkInsightsAccessScopeCmd.MarkFlagRequired("network-insights-access-scope-id")
	ec2_deleteNetworkInsightsAccessScopeCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_deleteNetworkInsightsAccessScopeCmd)
}

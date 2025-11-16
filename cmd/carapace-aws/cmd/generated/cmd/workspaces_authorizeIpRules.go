package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_authorizeIpRulesCmd = &cobra.Command{
	Use:   "authorize-ip-rules",
	Short: "Adds one or more rules to the specified IP access control group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_authorizeIpRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_authorizeIpRulesCmd).Standalone()

		workspaces_authorizeIpRulesCmd.Flags().String("group-id", "", "The identifier of the group.")
		workspaces_authorizeIpRulesCmd.Flags().String("user-rules", "", "The rules to add to the group.")
		workspaces_authorizeIpRulesCmd.MarkFlagRequired("group-id")
		workspaces_authorizeIpRulesCmd.MarkFlagRequired("user-rules")
	})
	workspacesCmd.AddCommand(workspaces_authorizeIpRulesCmd)
}

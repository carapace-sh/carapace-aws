package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_revokeIpRulesCmd = &cobra.Command{
	Use:   "revoke-ip-rules",
	Short: "Removes one or more rules from the specified IP access control group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_revokeIpRulesCmd).Standalone()

	workspaces_revokeIpRulesCmd.Flags().String("group-id", "", "The identifier of the group.")
	workspaces_revokeIpRulesCmd.Flags().String("user-rules", "", "The rules to remove from the group.")
	workspaces_revokeIpRulesCmd.MarkFlagRequired("group-id")
	workspaces_revokeIpRulesCmd.MarkFlagRequired("user-rules")
	workspacesCmd.AddCommand(workspaces_revokeIpRulesCmd)
}

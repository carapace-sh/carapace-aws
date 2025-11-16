package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_updateRulesOfIpGroupCmd = &cobra.Command{
	Use:   "update-rules-of-ip-group",
	Short: "Replaces the current rules of the specified IP access control group with the specified rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_updateRulesOfIpGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_updateRulesOfIpGroupCmd).Standalone()

		workspaces_updateRulesOfIpGroupCmd.Flags().String("group-id", "", "The identifier of the group.")
		workspaces_updateRulesOfIpGroupCmd.Flags().String("user-rules", "", "One or more rules.")
		workspaces_updateRulesOfIpGroupCmd.MarkFlagRequired("group-id")
		workspaces_updateRulesOfIpGroupCmd.MarkFlagRequired("user-rules")
	})
	workspacesCmd.AddCommand(workspaces_updateRulesOfIpGroupCmd)
}

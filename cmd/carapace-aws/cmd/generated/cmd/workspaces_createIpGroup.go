package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_createIpGroupCmd = &cobra.Command{
	Use:   "create-ip-group",
	Short: "Creates an IP access control group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_createIpGroupCmd).Standalone()

	workspaces_createIpGroupCmd.Flags().String("group-desc", "", "The description of the group.")
	workspaces_createIpGroupCmd.Flags().String("group-name", "", "The name of the group.")
	workspaces_createIpGroupCmd.Flags().String("tags", "", "The tags.")
	workspaces_createIpGroupCmd.Flags().String("user-rules", "", "The rules to add to the group.")
	workspaces_createIpGroupCmd.MarkFlagRequired("group-name")
	workspacesCmd.AddCommand(workspaces_createIpGroupCmd)
}

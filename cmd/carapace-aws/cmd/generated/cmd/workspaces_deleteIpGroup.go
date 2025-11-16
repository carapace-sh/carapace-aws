package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_deleteIpGroupCmd = &cobra.Command{
	Use:   "delete-ip-group",
	Short: "Deletes the specified IP access control group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_deleteIpGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_deleteIpGroupCmd).Standalone()

		workspaces_deleteIpGroupCmd.Flags().String("group-id", "", "The identifier of the IP access control group.")
		workspaces_deleteIpGroupCmd.MarkFlagRequired("group-id")
	})
	workspacesCmd.AddCommand(workspaces_deleteIpGroupCmd)
}

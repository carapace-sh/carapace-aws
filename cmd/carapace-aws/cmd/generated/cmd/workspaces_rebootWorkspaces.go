package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_rebootWorkspacesCmd = &cobra.Command{
	Use:   "reboot-workspaces",
	Short: "Reboots the specified WorkSpaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_rebootWorkspacesCmd).Standalone()

	workspaces_rebootWorkspacesCmd.Flags().String("reboot-workspace-requests", "", "The WorkSpaces to reboot.")
	workspaces_rebootWorkspacesCmd.MarkFlagRequired("reboot-workspace-requests")
	workspacesCmd.AddCommand(workspaces_rebootWorkspacesCmd)
}

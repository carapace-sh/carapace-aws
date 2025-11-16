package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_rebuildWorkspacesCmd = &cobra.Command{
	Use:   "rebuild-workspaces",
	Short: "Rebuilds the specified WorkSpace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_rebuildWorkspacesCmd).Standalone()

	workspaces_rebuildWorkspacesCmd.Flags().String("rebuild-workspace-requests", "", "The WorkSpace to rebuild.")
	workspaces_rebuildWorkspacesCmd.MarkFlagRequired("rebuild-workspace-requests")
	workspacesCmd.AddCommand(workspaces_rebuildWorkspacesCmd)
}

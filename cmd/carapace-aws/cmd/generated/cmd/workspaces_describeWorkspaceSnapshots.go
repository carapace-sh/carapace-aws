package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeWorkspaceSnapshotsCmd = &cobra.Command{
	Use:   "describe-workspace-snapshots",
	Short: "Describes the snapshots for the specified WorkSpace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeWorkspaceSnapshotsCmd).Standalone()

	workspaces_describeWorkspaceSnapshotsCmd.Flags().String("workspace-id", "", "The identifier of the WorkSpace.")
	workspaces_describeWorkspaceSnapshotsCmd.MarkFlagRequired("workspace-id")
	workspacesCmd.AddCommand(workspaces_describeWorkspaceSnapshotsCmd)
}

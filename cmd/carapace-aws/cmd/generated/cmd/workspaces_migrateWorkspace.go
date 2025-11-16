package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_migrateWorkspaceCmd = &cobra.Command{
	Use:   "migrate-workspace",
	Short: "Migrates a WorkSpace from one operating system or bundle type to another, while retaining the data on the user volume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_migrateWorkspaceCmd).Standalone()

	workspaces_migrateWorkspaceCmd.Flags().String("bundle-id", "", "The identifier of the target bundle type to migrate the WorkSpace to.")
	workspaces_migrateWorkspaceCmd.Flags().String("source-workspace-id", "", "The identifier of the WorkSpace to migrate from.")
	workspaces_migrateWorkspaceCmd.MarkFlagRequired("bundle-id")
	workspaces_migrateWorkspaceCmd.MarkFlagRequired("source-workspace-id")
	workspacesCmd.AddCommand(workspaces_migrateWorkspaceCmd)
}

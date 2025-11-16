package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_createStandbyWorkspacesCmd = &cobra.Command{
	Use:   "create-standby-workspaces",
	Short: "Creates a standby WorkSpace in a secondary Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_createStandbyWorkspacesCmd).Standalone()

	workspaces_createStandbyWorkspacesCmd.Flags().String("primary-region", "", "The Region of the primary WorkSpace.")
	workspaces_createStandbyWorkspacesCmd.Flags().String("standby-workspaces", "", "Information about the standby WorkSpace to be created.")
	workspaces_createStandbyWorkspacesCmd.MarkFlagRequired("primary-region")
	workspaces_createStandbyWorkspacesCmd.MarkFlagRequired("standby-workspaces")
	workspacesCmd.AddCommand(workspaces_createStandbyWorkspacesCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_deleteWorkspaceCmd = &cobra.Command{
	Use:   "delete-workspace",
	Short: "Deletes a workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_deleteWorkspaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_deleteWorkspaceCmd).Standalone()

		iottwinmaker_deleteWorkspaceCmd.Flags().String("workspace-id", "", "The ID of the workspace to delete.")
		iottwinmaker_deleteWorkspaceCmd.MarkFlagRequired("workspace-id")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_deleteWorkspaceCmd)
}

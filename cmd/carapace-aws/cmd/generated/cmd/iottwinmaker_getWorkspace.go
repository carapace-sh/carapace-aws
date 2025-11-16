package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_getWorkspaceCmd = &cobra.Command{
	Use:   "get-workspace",
	Short: "Retrieves information about a workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_getWorkspaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_getWorkspaceCmd).Standalone()

		iottwinmaker_getWorkspaceCmd.Flags().String("workspace-id", "", "The ID of the workspace.")
		iottwinmaker_getWorkspaceCmd.MarkFlagRequired("workspace-id")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_getWorkspaceCmd)
}

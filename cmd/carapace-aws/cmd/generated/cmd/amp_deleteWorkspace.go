package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_deleteWorkspaceCmd = &cobra.Command{
	Use:   "delete-workspace",
	Short: "Deletes an existing workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_deleteWorkspaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_deleteWorkspaceCmd).Standalone()

		amp_deleteWorkspaceCmd.Flags().String("client-token", "", "A unique identifier that you can provide to ensure the idempotency of the request.")
		amp_deleteWorkspaceCmd.Flags().String("workspace-id", "", "The ID of the workspace to delete.")
		amp_deleteWorkspaceCmd.MarkFlagRequired("workspace-id")
	})
	ampCmd.AddCommand(amp_deleteWorkspaceCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_updateWorkspaceAliasCmd = &cobra.Command{
	Use:   "update-workspace-alias",
	Short: "Updates the alias of an existing workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_updateWorkspaceAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_updateWorkspaceAliasCmd).Standalone()

		amp_updateWorkspaceAliasCmd.Flags().String("alias", "", "The new alias for the workspace.")
		amp_updateWorkspaceAliasCmd.Flags().String("client-token", "", "A unique identifier that you can provide to ensure the idempotency of the request.")
		amp_updateWorkspaceAliasCmd.Flags().String("workspace-id", "", "The ID of the workspace to update.")
		amp_updateWorkspaceAliasCmd.MarkFlagRequired("workspace-id")
	})
	ampCmd.AddCommand(amp_updateWorkspaceAliasCmd)
}

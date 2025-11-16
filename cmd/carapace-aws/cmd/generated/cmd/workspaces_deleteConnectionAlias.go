package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_deleteConnectionAliasCmd = &cobra.Command{
	Use:   "delete-connection-alias",
	Short: "Deletes the specified connection alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_deleteConnectionAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_deleteConnectionAliasCmd).Standalone()

		workspaces_deleteConnectionAliasCmd.Flags().String("alias-id", "", "The identifier of the connection alias to delete.")
		workspaces_deleteConnectionAliasCmd.MarkFlagRequired("alias-id")
	})
	workspacesCmd.AddCommand(workspaces_deleteConnectionAliasCmd)
}

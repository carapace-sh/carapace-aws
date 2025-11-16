package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_updateConnectionAliasPermissionCmd = &cobra.Command{
	Use:   "update-connection-alias-permission",
	Short: "Shares or unshares a connection alias with one account by specifying whether that account has permission to associate the connection alias with a directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_updateConnectionAliasPermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_updateConnectionAliasPermissionCmd).Standalone()

		workspaces_updateConnectionAliasPermissionCmd.Flags().String("alias-id", "", "The identifier of the connection alias that you want to update permissions for.")
		workspaces_updateConnectionAliasPermissionCmd.Flags().String("connection-alias-permission", "", "Indicates whether to share or unshare the connection alias with the specified Amazon Web Services account.")
		workspaces_updateConnectionAliasPermissionCmd.MarkFlagRequired("alias-id")
		workspaces_updateConnectionAliasPermissionCmd.MarkFlagRequired("connection-alias-permission")
	})
	workspacesCmd.AddCommand(workspaces_updateConnectionAliasPermissionCmd)
}

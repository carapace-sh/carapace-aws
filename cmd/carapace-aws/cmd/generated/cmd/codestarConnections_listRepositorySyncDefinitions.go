package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_listRepositorySyncDefinitionsCmd = &cobra.Command{
	Use:   "list-repository-sync-definitions",
	Short: "Lists the repository sync definitions for repository links in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_listRepositorySyncDefinitionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codestarConnections_listRepositorySyncDefinitionsCmd).Standalone()

		codestarConnections_listRepositorySyncDefinitionsCmd.Flags().String("repository-link-id", "", "The ID of the repository link for the sync definition for which you want to retrieve information.")
		codestarConnections_listRepositorySyncDefinitionsCmd.Flags().String("sync-type", "", "The sync type of the repository link for the the sync definition for which you want to retrieve information.")
		codestarConnections_listRepositorySyncDefinitionsCmd.MarkFlagRequired("repository-link-id")
		codestarConnections_listRepositorySyncDefinitionsCmd.MarkFlagRequired("sync-type")
	})
	codestarConnectionsCmd.AddCommand(codestarConnections_listRepositorySyncDefinitionsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_listRepositorySyncDefinitionsCmd = &cobra.Command{
	Use:   "list-repository-sync-definitions",
	Short: "Lists the repository sync definitions for repository links in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_listRepositorySyncDefinitionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeconnections_listRepositorySyncDefinitionsCmd).Standalone()

		codeconnections_listRepositorySyncDefinitionsCmd.Flags().String("repository-link-id", "", "The ID of the repository link for the sync definition for which you want to retrieve information.")
		codeconnections_listRepositorySyncDefinitionsCmd.Flags().String("sync-type", "", "The sync type of the repository link for the the sync definition for which you want to retrieve information.")
		codeconnections_listRepositorySyncDefinitionsCmd.MarkFlagRequired("repository-link-id")
		codeconnections_listRepositorySyncDefinitionsCmd.MarkFlagRequired("sync-type")
	})
	codeconnectionsCmd.AddCommand(codeconnections_listRepositorySyncDefinitionsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_listRepositorySyncDefinitionsCmd = &cobra.Command{
	Use:   "list-repository-sync-definitions",
	Short: "List repository sync definitions with detail data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_listRepositorySyncDefinitionsCmd).Standalone()

	proton_listRepositorySyncDefinitionsCmd.Flags().String("next-token", "", "A token that indicates the location of the next repository sync definition in the array of repository sync definitions, after the list of repository sync definitions previously requested.")
	proton_listRepositorySyncDefinitionsCmd.Flags().String("repository-name", "", "The repository name.")
	proton_listRepositorySyncDefinitionsCmd.Flags().String("repository-provider", "", "The repository provider.")
	proton_listRepositorySyncDefinitionsCmd.Flags().String("sync-type", "", "The sync type.")
	proton_listRepositorySyncDefinitionsCmd.MarkFlagRequired("repository-name")
	proton_listRepositorySyncDefinitionsCmd.MarkFlagRequired("repository-provider")
	proton_listRepositorySyncDefinitionsCmd.MarkFlagRequired("sync-type")
	protonCmd.AddCommand(proton_listRepositorySyncDefinitionsCmd)
}

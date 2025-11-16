package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_getRepositoryCatalogDataCmd = &cobra.Command{
	Use:   "get-repository-catalog-data",
	Short: "Retrieve catalog metadata for a repository in a public registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_getRepositoryCatalogDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecrPublic_getRepositoryCatalogDataCmd).Standalone()

		ecrPublic_getRepositoryCatalogDataCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID that's associated with the registry that contains the repositories to be described.")
		ecrPublic_getRepositoryCatalogDataCmd.Flags().String("repository-name", "", "The name of the repository to retrieve the catalog metadata for.")
		ecrPublic_getRepositoryCatalogDataCmd.MarkFlagRequired("repository-name")
	})
	ecrPublicCmd.AddCommand(ecrPublic_getRepositoryCatalogDataCmd)
}

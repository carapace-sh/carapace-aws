package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_putRepositoryCatalogDataCmd = &cobra.Command{
	Use:   "put-repository-catalog-data",
	Short: "Creates or updates the catalog data for a repository in a public registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_putRepositoryCatalogDataCmd).Standalone()

	ecrPublic_putRepositoryCatalogDataCmd.Flags().String("catalog-data", "", "An object containing the catalog data for a repository.")
	ecrPublic_putRepositoryCatalogDataCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID that's associated with the public registry the repository is in.")
	ecrPublic_putRepositoryCatalogDataCmd.Flags().String("repository-name", "", "The name of the repository to create or update the catalog data for.")
	ecrPublic_putRepositoryCatalogDataCmd.MarkFlagRequired("catalog-data")
	ecrPublic_putRepositoryCatalogDataCmd.MarkFlagRequired("repository-name")
	ecrPublicCmd.AddCommand(ecrPublic_putRepositoryCatalogDataCmd)
}

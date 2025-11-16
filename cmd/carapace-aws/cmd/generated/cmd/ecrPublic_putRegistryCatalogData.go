package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_putRegistryCatalogDataCmd = &cobra.Command{
	Use:   "put-registry-catalog-data",
	Short: "Create or update the catalog data for a public registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_putRegistryCatalogDataCmd).Standalone()

	ecrPublic_putRegistryCatalogDataCmd.Flags().String("display-name", "", "The display name for a public registry.")
	ecrPublicCmd.AddCommand(ecrPublic_putRegistryCatalogDataCmd)
}

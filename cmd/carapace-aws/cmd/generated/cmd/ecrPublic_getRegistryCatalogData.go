package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_getRegistryCatalogDataCmd = &cobra.Command{
	Use:   "get-registry-catalog-data",
	Short: "Retrieves catalog metadata for a public registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_getRegistryCatalogDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecrPublic_getRegistryCatalogDataCmd).Standalone()

	})
	ecrPublicCmd.AddCommand(ecrPublic_getRegistryCatalogDataCmd)
}

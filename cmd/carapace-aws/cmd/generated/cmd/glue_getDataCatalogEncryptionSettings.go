package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getDataCatalogEncryptionSettingsCmd = &cobra.Command{
	Use:   "get-data-catalog-encryption-settings",
	Short: "Retrieves the security configuration for a specified catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getDataCatalogEncryptionSettingsCmd).Standalone()

	glue_getDataCatalogEncryptionSettingsCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog to retrieve the security configuration for.")
	glueCmd.AddCommand(glue_getDataCatalogEncryptionSettingsCmd)
}

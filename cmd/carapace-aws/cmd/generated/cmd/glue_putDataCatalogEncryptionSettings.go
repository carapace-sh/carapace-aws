package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_putDataCatalogEncryptionSettingsCmd = &cobra.Command{
	Use:   "put-data-catalog-encryption-settings",
	Short: "Sets the security configuration for a specified catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_putDataCatalogEncryptionSettingsCmd).Standalone()

	glue_putDataCatalogEncryptionSettingsCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog to set the security configuration for.")
	glue_putDataCatalogEncryptionSettingsCmd.Flags().String("data-catalog-encryption-settings", "", "The security configuration to set.")
	glue_putDataCatalogEncryptionSettingsCmd.MarkFlagRequired("data-catalog-encryption-settings")
	glueCmd.AddCommand(glue_putDataCatalogEncryptionSettingsCmd)
}

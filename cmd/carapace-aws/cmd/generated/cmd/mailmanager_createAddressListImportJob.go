package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_createAddressListImportJobCmd = &cobra.Command{
	Use:   "create-address-list-import-job",
	Short: "Creates an import job for an address list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_createAddressListImportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_createAddressListImportJobCmd).Standalone()

		mailmanager_createAddressListImportJobCmd.Flags().String("address-list-id", "", "The unique identifier of the address list for importing addresses to.")
		mailmanager_createAddressListImportJobCmd.Flags().String("client-token", "", "A unique token that Amazon SES uses to recognize subsequent retries of the same request.")
		mailmanager_createAddressListImportJobCmd.Flags().String("import-data-format", "", "The format of the input for an import job.")
		mailmanager_createAddressListImportJobCmd.Flags().String("name", "", "A user-friendly name for the import job.")
		mailmanager_createAddressListImportJobCmd.MarkFlagRequired("address-list-id")
		mailmanager_createAddressListImportJobCmd.MarkFlagRequired("import-data-format")
		mailmanager_createAddressListImportJobCmd.MarkFlagRequired("name")
	})
	mailmanagerCmd.AddCommand(mailmanager_createAddressListImportJobCmd)
}

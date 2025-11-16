package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_getAddressListImportJobCmd = &cobra.Command{
	Use:   "get-address-list-import-job",
	Short: "Fetch attributes of an import job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_getAddressListImportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_getAddressListImportJobCmd).Standalone()

		mailmanager_getAddressListImportJobCmd.Flags().String("job-id", "", "The identifier of the import job that needs to be retrieved.")
		mailmanager_getAddressListImportJobCmd.MarkFlagRequired("job-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_getAddressListImportJobCmd)
}

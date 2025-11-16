package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_stopAddressListImportJobCmd = &cobra.Command{
	Use:   "stop-address-list-import-job",
	Short: "Stops an ongoing import job for an address list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_stopAddressListImportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_stopAddressListImportJobCmd).Standalone()

		mailmanager_stopAddressListImportJobCmd.Flags().String("job-id", "", "The identifier of the import job that needs to be stopped.")
		mailmanager_stopAddressListImportJobCmd.MarkFlagRequired("job-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_stopAddressListImportJobCmd)
}

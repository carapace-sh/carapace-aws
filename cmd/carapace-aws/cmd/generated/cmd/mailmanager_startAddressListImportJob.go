package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_startAddressListImportJobCmd = &cobra.Command{
	Use:   "start-address-list-import-job",
	Short: "Starts an import job for an address list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_startAddressListImportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_startAddressListImportJobCmd).Standalone()

		mailmanager_startAddressListImportJobCmd.Flags().String("job-id", "", "The identifier of the import job that needs to be started.")
		mailmanager_startAddressListImportJobCmd.MarkFlagRequired("job-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_startAddressListImportJobCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_listAddressListImportJobsCmd = &cobra.Command{
	Use:   "list-address-list-import-jobs",
	Short: "Lists jobs for an address list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_listAddressListImportJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_listAddressListImportJobsCmd).Standalone()

		mailmanager_listAddressListImportJobsCmd.Flags().String("address-list-id", "", "The unique identifier of the address list for listing import jobs.")
		mailmanager_listAddressListImportJobsCmd.Flags().String("next-token", "", "If you received a pagination token from a previous call to this API, you can provide it here to continue paginating through the next page of results.")
		mailmanager_listAddressListImportJobsCmd.Flags().String("page-size", "", "The maximum number of import jobs that are returned per call.")
		mailmanager_listAddressListImportJobsCmd.MarkFlagRequired("address-list-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_listAddressListImportJobsCmd)
}

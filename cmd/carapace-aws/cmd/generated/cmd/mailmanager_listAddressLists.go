package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_listAddressListsCmd = &cobra.Command{
	Use:   "list-address-lists",
	Short: "Lists address lists for this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_listAddressListsCmd).Standalone()

	mailmanager_listAddressListsCmd.Flags().String("next-token", "", "If you received a pagination token from a previous call to this API, you can provide it here to continue paginating through the next page of results.")
	mailmanager_listAddressListsCmd.Flags().String("page-size", "", "The maximum number of address list resources that are returned per call.")
	mailmanagerCmd.AddCommand(mailmanager_listAddressListsCmd)
}

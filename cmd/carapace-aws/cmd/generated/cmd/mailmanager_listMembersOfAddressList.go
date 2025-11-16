package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_listMembersOfAddressListCmd = &cobra.Command{
	Use:   "list-members-of-address-list",
	Short: "Lists members of an address list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_listMembersOfAddressListCmd).Standalone()

	mailmanager_listMembersOfAddressListCmd.Flags().String("address-list-id", "", "The unique identifier of the address list to list the addresses from.")
	mailmanager_listMembersOfAddressListCmd.Flags().String("filter", "", "Filter to be used to limit the results.")
	mailmanager_listMembersOfAddressListCmd.Flags().String("next-token", "", "If you received a pagination token from a previous call to this API, you can provide it here to continue paginating through the next page of results.")
	mailmanager_listMembersOfAddressListCmd.Flags().String("page-size", "", "The maximum number of address list members that are returned per call.")
	mailmanager_listMembersOfAddressListCmd.MarkFlagRequired("address-list-id")
	mailmanagerCmd.AddCommand(mailmanager_listMembersOfAddressListCmd)
}

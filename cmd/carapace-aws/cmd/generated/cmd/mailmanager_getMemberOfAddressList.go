package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_getMemberOfAddressListCmd = &cobra.Command{
	Use:   "get-member-of-address-list",
	Short: "Fetch attributes of a member in an address list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_getMemberOfAddressListCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_getMemberOfAddressListCmd).Standalone()

		mailmanager_getMemberOfAddressListCmd.Flags().String("address", "", "The address to be retrieved from the address list.")
		mailmanager_getMemberOfAddressListCmd.Flags().String("address-list-id", "", "The unique identifier of the address list to retrieve the address from.")
		mailmanager_getMemberOfAddressListCmd.MarkFlagRequired("address")
		mailmanager_getMemberOfAddressListCmd.MarkFlagRequired("address-list-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_getMemberOfAddressListCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_deregisterMemberFromAddressListCmd = &cobra.Command{
	Use:   "deregister-member-from-address-list",
	Short: "Removes a member from an address list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_deregisterMemberFromAddressListCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_deregisterMemberFromAddressListCmd).Standalone()

		mailmanager_deregisterMemberFromAddressListCmd.Flags().String("address", "", "The address to be removed from the address list.")
		mailmanager_deregisterMemberFromAddressListCmd.Flags().String("address-list-id", "", "The unique identifier of the address list to remove the address from.")
		mailmanager_deregisterMemberFromAddressListCmd.MarkFlagRequired("address")
		mailmanager_deregisterMemberFromAddressListCmd.MarkFlagRequired("address-list-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_deregisterMemberFromAddressListCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_registerMemberToAddressListCmd = &cobra.Command{
	Use:   "register-member-to-address-list",
	Short: "Adds a member to an address list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_registerMemberToAddressListCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_registerMemberToAddressListCmd).Standalone()

		mailmanager_registerMemberToAddressListCmd.Flags().String("address", "", "The address to be added to the address list.")
		mailmanager_registerMemberToAddressListCmd.Flags().String("address-list-id", "", "The unique identifier of the address list where the address should be added.")
		mailmanager_registerMemberToAddressListCmd.MarkFlagRequired("address")
		mailmanager_registerMemberToAddressListCmd.MarkFlagRequired("address-list-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_registerMemberToAddressListCmd)
}

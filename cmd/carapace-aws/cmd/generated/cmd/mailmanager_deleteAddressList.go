package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_deleteAddressListCmd = &cobra.Command{
	Use:   "delete-address-list",
	Short: "Deletes an address list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_deleteAddressListCmd).Standalone()

	mailmanager_deleteAddressListCmd.Flags().String("address-list-id", "", "The identifier of an existing address list resource to delete.")
	mailmanager_deleteAddressListCmd.MarkFlagRequired("address-list-id")
	mailmanagerCmd.AddCommand(mailmanager_deleteAddressListCmd)
}

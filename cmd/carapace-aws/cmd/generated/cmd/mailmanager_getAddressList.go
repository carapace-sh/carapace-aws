package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_getAddressListCmd = &cobra.Command{
	Use:   "get-address-list",
	Short: "Fetch attributes of an address list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_getAddressListCmd).Standalone()

	mailmanager_getAddressListCmd.Flags().String("address-list-id", "", "The identifier of an existing address list resource to be retrieved.")
	mailmanager_getAddressListCmd.MarkFlagRequired("address-list-id")
	mailmanagerCmd.AddCommand(mailmanager_getAddressListCmd)
}

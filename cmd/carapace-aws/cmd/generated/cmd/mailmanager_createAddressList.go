package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_createAddressListCmd = &cobra.Command{
	Use:   "create-address-list",
	Short: "Creates a new address list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_createAddressListCmd).Standalone()

	mailmanager_createAddressListCmd.Flags().String("address-list-name", "", "A user-friendly name for the address list.")
	mailmanager_createAddressListCmd.Flags().String("client-token", "", "A unique token that Amazon SES uses to recognize subsequent retries of the same request.")
	mailmanager_createAddressListCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for the resource.")
	mailmanager_createAddressListCmd.MarkFlagRequired("address-list-name")
	mailmanagerCmd.AddCommand(mailmanager_createAddressListCmd)
}

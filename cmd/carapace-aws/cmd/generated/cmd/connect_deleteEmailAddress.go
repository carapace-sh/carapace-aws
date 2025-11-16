package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteEmailAddressCmd = &cobra.Command{
	Use:   "delete-email-address",
	Short: "Deletes email address from the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteEmailAddressCmd).Standalone()

	connect_deleteEmailAddressCmd.Flags().String("email-address-id", "", "The identifier of the email address.")
	connect_deleteEmailAddressCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_deleteEmailAddressCmd.MarkFlagRequired("email-address-id")
	connect_deleteEmailAddressCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_deleteEmailAddressCmd)
}

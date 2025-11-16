package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createEmailAddressCmd = &cobra.Command{
	Use:   "create-email-address",
	Short: "Create new email address in the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createEmailAddressCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_createEmailAddressCmd).Standalone()

		connect_createEmailAddressCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_createEmailAddressCmd.Flags().String("description", "", "The description of the email address.")
		connect_createEmailAddressCmd.Flags().String("display-name", "", "The display name of email address")
		connect_createEmailAddressCmd.Flags().String("email-address", "", "The email address, including the domain.")
		connect_createEmailAddressCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_createEmailAddressCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		connect_createEmailAddressCmd.MarkFlagRequired("email-address")
		connect_createEmailAddressCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_createEmailAddressCmd)
}

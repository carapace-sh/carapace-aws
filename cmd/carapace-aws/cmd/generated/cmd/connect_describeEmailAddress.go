package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeEmailAddressCmd = &cobra.Command{
	Use:   "describe-email-address",
	Short: "Describe email address form the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeEmailAddressCmd).Standalone()

	connect_describeEmailAddressCmd.Flags().String("email-address-id", "", "The identifier of the email address.")
	connect_describeEmailAddressCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_describeEmailAddressCmd.MarkFlagRequired("email-address-id")
	connect_describeEmailAddressCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_describeEmailAddressCmd)
}

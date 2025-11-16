package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_getPhoneNumberOrderCmd = &cobra.Command{
	Use:   "get-phone-number-order",
	Short: "Retrieves details for the specified phone number order, such as the order creation timestamp, phone numbers in E.164 format, product type, and order status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_getPhoneNumberOrderCmd).Standalone()

	chime_getPhoneNumberOrderCmd.Flags().String("phone-number-order-id", "", "The ID for the phone number order.")
	chime_getPhoneNumberOrderCmd.MarkFlagRequired("phone-number-order-id")
	chimeCmd.AddCommand(chime_getPhoneNumberOrderCmd)
}

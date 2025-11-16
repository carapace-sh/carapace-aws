package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_createPhoneNumberOrderCmd = &cobra.Command{
	Use:   "create-phone-number-order",
	Short: "Creates an order for phone numbers to be provisioned.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_createPhoneNumberOrderCmd).Standalone()

	chime_createPhoneNumberOrderCmd.Flags().String("e164-phone-numbers", "", "List of phone numbers, in E.164 format.")
	chime_createPhoneNumberOrderCmd.Flags().String("product-type", "", "The phone number product type.")
	chime_createPhoneNumberOrderCmd.MarkFlagRequired("e164-phone-numbers")
	chime_createPhoneNumberOrderCmd.MarkFlagRequired("product-type")
	chimeCmd.AddCommand(chime_createPhoneNumberOrderCmd)
}

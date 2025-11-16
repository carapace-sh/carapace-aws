package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_updatePhoneNumberCmd = &cobra.Command{
	Use:   "update-phone-number",
	Short: "Updates phone number details, such as product type or calling name, for the specified phone number ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_updatePhoneNumberCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_updatePhoneNumberCmd).Standalone()

		chime_updatePhoneNumberCmd.Flags().String("calling-name", "", "The outbound calling name associated with the phone number.")
		chime_updatePhoneNumberCmd.Flags().String("phone-number-id", "", "The phone number ID.")
		chime_updatePhoneNumberCmd.Flags().String("product-type", "", "The product type.")
		chime_updatePhoneNumberCmd.MarkFlagRequired("phone-number-id")
	})
	chimeCmd.AddCommand(chime_updatePhoneNumberCmd)
}

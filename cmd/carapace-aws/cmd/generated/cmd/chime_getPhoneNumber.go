package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_getPhoneNumberCmd = &cobra.Command{
	Use:   "get-phone-number",
	Short: "Retrieves details for the specified phone number ID, such as associations, capabilities, and product type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_getPhoneNumberCmd).Standalone()

	chime_getPhoneNumberCmd.Flags().String("phone-number-id", "", "The phone number ID.")
	chime_getPhoneNumberCmd.MarkFlagRequired("phone-number-id")
	chimeCmd.AddCommand(chime_getPhoneNumberCmd)
}

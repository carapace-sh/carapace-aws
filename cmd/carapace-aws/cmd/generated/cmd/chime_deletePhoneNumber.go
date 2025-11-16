package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_deletePhoneNumberCmd = &cobra.Command{
	Use:   "delete-phone-number",
	Short: "Moves the specified phone number into the **Deletion queue**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_deletePhoneNumberCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_deletePhoneNumberCmd).Standalone()

		chime_deletePhoneNumberCmd.Flags().String("phone-number-id", "", "The phone number ID.")
		chime_deletePhoneNumberCmd.MarkFlagRequired("phone-number-id")
	})
	chimeCmd.AddCommand(chime_deletePhoneNumberCmd)
}

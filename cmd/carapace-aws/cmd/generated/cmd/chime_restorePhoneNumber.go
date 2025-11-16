package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_restorePhoneNumberCmd = &cobra.Command{
	Use:   "restore-phone-number",
	Short: "Moves a phone number from the **Deletion queue** back into the phone number **Inventory**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_restorePhoneNumberCmd).Standalone()

	chime_restorePhoneNumberCmd.Flags().String("phone-number-id", "", "The phone number.")
	chime_restorePhoneNumberCmd.MarkFlagRequired("phone-number-id")
	chimeCmd.AddCommand(chime_restorePhoneNumberCmd)
}

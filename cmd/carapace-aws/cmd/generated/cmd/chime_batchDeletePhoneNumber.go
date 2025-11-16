package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_batchDeletePhoneNumberCmd = &cobra.Command{
	Use:   "batch-delete-phone-number",
	Short: "Moves phone numbers into the **Deletion queue**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_batchDeletePhoneNumberCmd).Standalone()

	chime_batchDeletePhoneNumberCmd.Flags().String("phone-number-ids", "", "List of phone number IDs.")
	chime_batchDeletePhoneNumberCmd.MarkFlagRequired("phone-number-ids")
	chimeCmd.AddCommand(chime_batchDeletePhoneNumberCmd)
}

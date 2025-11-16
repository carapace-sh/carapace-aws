package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_batchUpdatePhoneNumberCmd = &cobra.Command{
	Use:   "batch-update-phone-number",
	Short: "Updates phone number product types or calling names.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_batchUpdatePhoneNumberCmd).Standalone()

	chime_batchUpdatePhoneNumberCmd.Flags().String("update-phone-number-request-items", "", "The request containing the phone number IDs and product types or calling names to update.")
	chime_batchUpdatePhoneNumberCmd.MarkFlagRequired("update-phone-number-request-items")
	chimeCmd.AddCommand(chime_batchUpdatePhoneNumberCmd)
}

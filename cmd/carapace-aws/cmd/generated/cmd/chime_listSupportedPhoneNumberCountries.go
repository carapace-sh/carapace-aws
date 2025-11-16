package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_listSupportedPhoneNumberCountriesCmd = &cobra.Command{
	Use:   "list-supported-phone-number-countries",
	Short: "Lists supported phone number countries.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_listSupportedPhoneNumberCountriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_listSupportedPhoneNumberCountriesCmd).Standalone()

		chime_listSupportedPhoneNumberCountriesCmd.Flags().String("product-type", "", "The phone number product type.")
		chime_listSupportedPhoneNumberCountriesCmd.MarkFlagRequired("product-type")
	})
	chimeCmd.AddCommand(chime_listSupportedPhoneNumberCountriesCmd)
}

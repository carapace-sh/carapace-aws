package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_listPhoneNumberOrdersCmd = &cobra.Command{
	Use:   "list-phone-number-orders",
	Short: "Lists the phone number orders for the administrator's Amazon Chime account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_listPhoneNumberOrdersCmd).Standalone()

	chime_listPhoneNumberOrdersCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	chime_listPhoneNumberOrdersCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
	chimeCmd.AddCommand(chime_listPhoneNumberOrdersCmd)
}

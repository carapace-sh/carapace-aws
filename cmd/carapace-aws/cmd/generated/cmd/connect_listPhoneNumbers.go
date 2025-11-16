package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listPhoneNumbersCmd = &cobra.Command{
	Use:   "list-phone-numbers",
	Short: "Provides information about the phone numbers for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listPhoneNumbersCmd).Standalone()

	connect_listPhoneNumbersCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_listPhoneNumbersCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_listPhoneNumbersCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_listPhoneNumbersCmd.Flags().String("phone-number-country-codes", "", "The ISO country code.")
	connect_listPhoneNumbersCmd.Flags().String("phone-number-types", "", "The type of phone number.")
	connect_listPhoneNumbersCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_listPhoneNumbersCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_searchAvailablePhoneNumbersCmd = &cobra.Command{
	Use:   "search-available-phone-numbers",
	Short: "Searches for available phone numbers that you can claim to your Amazon Connect instance or traffic distribution group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_searchAvailablePhoneNumbersCmd).Standalone()

	connect_searchAvailablePhoneNumbersCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance that phone numbers are claimed to.")
	connect_searchAvailablePhoneNumbersCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_searchAvailablePhoneNumbersCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_searchAvailablePhoneNumbersCmd.Flags().String("phone-number-country-code", "", "The ISO country code.")
	connect_searchAvailablePhoneNumbersCmd.Flags().String("phone-number-prefix", "", "The prefix of the phone number.")
	connect_searchAvailablePhoneNumbersCmd.Flags().String("phone-number-type", "", "The type of phone number.")
	connect_searchAvailablePhoneNumbersCmd.Flags().String("target-arn", "", "The Amazon Resource Name (ARN) for Amazon Connect instances or traffic distribution groups that phone number inbound traffic is routed through.")
	connect_searchAvailablePhoneNumbersCmd.MarkFlagRequired("phone-number-country-code")
	connect_searchAvailablePhoneNumbersCmd.MarkFlagRequired("phone-number-type")
	connectCmd.AddCommand(connect_searchAvailablePhoneNumbersCmd)
}

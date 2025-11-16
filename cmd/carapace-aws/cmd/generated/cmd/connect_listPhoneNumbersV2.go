package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listPhoneNumbersV2Cmd = &cobra.Command{
	Use:   "list-phone-numbers-v2",
	Short: "Lists phone numbers claimed to your Amazon Connect instance or traffic distribution group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listPhoneNumbersV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listPhoneNumbersV2Cmd).Standalone()

		connect_listPhoneNumbersV2Cmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance that phone numbers are claimed to.")
		connect_listPhoneNumbersV2Cmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listPhoneNumbersV2Cmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listPhoneNumbersV2Cmd.Flags().String("phone-number-country-codes", "", "The ISO country code.")
		connect_listPhoneNumbersV2Cmd.Flags().String("phone-number-prefix", "", "The prefix of the phone number.")
		connect_listPhoneNumbersV2Cmd.Flags().String("phone-number-types", "", "The type of phone number.")
		connect_listPhoneNumbersV2Cmd.Flags().String("target-arn", "", "The Amazon Resource Name (ARN) for Amazon Connect instances or traffic distribution groups that phone number inbound traffic is routed through.")
	})
	connectCmd.AddCommand(connect_listPhoneNumbersV2Cmd)
}

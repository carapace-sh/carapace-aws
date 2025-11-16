package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_listPhoneNumbersOptedOutCmd = &cobra.Command{
	Use:   "list-phone-numbers-opted-out",
	Short: "Returns a list of phone numbers that are opted out, meaning you cannot send SMS messages to them.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_listPhoneNumbersOptedOutCmd).Standalone()

	sns_listPhoneNumbersOptedOutCmd.Flags().String("next-token", "", "A `NextToken` string is used when you call the `ListPhoneNumbersOptedOut` action to retrieve additional records that are available after the first page of results.")
	snsCmd.AddCommand(sns_listPhoneNumbersOptedOutCmd)
}

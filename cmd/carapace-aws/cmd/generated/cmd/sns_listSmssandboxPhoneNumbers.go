package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_listSmssandboxPhoneNumbersCmd = &cobra.Command{
	Use:   "list-smssandbox-phone-numbers",
	Short: "Lists the calling Amazon Web Services account's current verified and pending destination phone numbers in the SMS sandbox.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_listSmssandboxPhoneNumbersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_listSmssandboxPhoneNumbersCmd).Standalone()

		sns_listSmssandboxPhoneNumbersCmd.Flags().String("max-results", "", "The maximum number of phone numbers to return.")
		sns_listSmssandboxPhoneNumbersCmd.Flags().String("next-token", "", "Token that the previous `ListSMSSandboxPhoneNumbersInput` request returns.")
	})
	snsCmd.AddCommand(sns_listSmssandboxPhoneNumbersCmd)
}

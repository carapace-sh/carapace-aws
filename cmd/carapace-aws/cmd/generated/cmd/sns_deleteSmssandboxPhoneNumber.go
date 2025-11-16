package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_deleteSmssandboxPhoneNumberCmd = &cobra.Command{
	Use:   "delete-smssandbox-phone-number",
	Short: "Deletes an Amazon Web Services account's verified or pending phone number from the SMS sandbox.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_deleteSmssandboxPhoneNumberCmd).Standalone()

	sns_deleteSmssandboxPhoneNumberCmd.Flags().String("phone-number", "", "The destination phone number to delete.")
	sns_deleteSmssandboxPhoneNumberCmd.MarkFlagRequired("phone-number")
	snsCmd.AddCommand(sns_deleteSmssandboxPhoneNumberCmd)
}

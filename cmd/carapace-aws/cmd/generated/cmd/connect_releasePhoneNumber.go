package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_releasePhoneNumberCmd = &cobra.Command{
	Use:   "release-phone-number",
	Short: "Releases a phone number previously claimed to an Amazon Connect instance or traffic distribution group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_releasePhoneNumberCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_releasePhoneNumberCmd).Standalone()

		connect_releasePhoneNumberCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_releasePhoneNumberCmd.Flags().String("phone-number-id", "", "A unique identifier for the phone number.")
		connect_releasePhoneNumberCmd.MarkFlagRequired("phone-number-id")
	})
	connectCmd.AddCommand(connect_releasePhoneNumberCmd)
}

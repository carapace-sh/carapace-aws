package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_acceptPrimaryEmailUpdateCmd = &cobra.Command{
	Use:   "accept-primary-email-update",
	Short: "Accepts the request that originated from [StartPrimaryEmailUpdate]() to update the primary email address (also known as the root user email address) for the specified account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_acceptPrimaryEmailUpdateCmd).Standalone()

	account_acceptPrimaryEmailUpdateCmd.Flags().String("account-id", "", "Specifies the 12-digit account ID number of the Amazon Web Services account that you want to access or modify with this operation.")
	account_acceptPrimaryEmailUpdateCmd.Flags().String("otp", "", "The OTP code sent to the `PrimaryEmail` specified on the `StartPrimaryEmailUpdate` API call.")
	account_acceptPrimaryEmailUpdateCmd.Flags().String("primary-email", "", "The new primary email address for use with the specified account.")
	account_acceptPrimaryEmailUpdateCmd.MarkFlagRequired("account-id")
	account_acceptPrimaryEmailUpdateCmd.MarkFlagRequired("otp")
	account_acceptPrimaryEmailUpdateCmd.MarkFlagRequired("primary-email")
	accountCmd.AddCommand(account_acceptPrimaryEmailUpdateCmd)
}

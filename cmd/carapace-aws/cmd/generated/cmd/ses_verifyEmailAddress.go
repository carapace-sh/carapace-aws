package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_verifyEmailAddressCmd = &cobra.Command{
	Use:   "verify-email-address",
	Short: "Deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_verifyEmailAddressCmd).Standalone()

	ses_verifyEmailAddressCmd.Flags().String("email-address", "", "The email address to be verified.")
	ses_verifyEmailAddressCmd.MarkFlagRequired("email-address")
	sesCmd.AddCommand(ses_verifyEmailAddressCmd)
}

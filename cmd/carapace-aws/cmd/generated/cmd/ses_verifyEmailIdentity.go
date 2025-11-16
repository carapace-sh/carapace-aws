package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_verifyEmailIdentityCmd = &cobra.Command{
	Use:   "verify-email-identity",
	Short: "Adds an email address to the list of identities for your Amazon SES account in the current Amazon Web Services Region and attempts to verify it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_verifyEmailIdentityCmd).Standalone()

	ses_verifyEmailIdentityCmd.Flags().String("email-address", "", "The email address to be verified.")
	ses_verifyEmailIdentityCmd.MarkFlagRequired("email-address")
	sesCmd.AddCommand(ses_verifyEmailIdentityCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_sendCustomVerificationEmailCmd = &cobra.Command{
	Use:   "send-custom-verification-email",
	Short: "Adds an email address to the list of identities for your Amazon SES account in the current Amazon Web Services Region and attempts to verify it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_sendCustomVerificationEmailCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_sendCustomVerificationEmailCmd).Standalone()

		ses_sendCustomVerificationEmailCmd.Flags().String("configuration-set-name", "", "Name of a configuration set to use when sending the verification email.")
		ses_sendCustomVerificationEmailCmd.Flags().String("email-address", "", "The email address to verify.")
		ses_sendCustomVerificationEmailCmd.Flags().String("template-name", "", "The name of the custom verification email template to use when sending the verification email.")
		ses_sendCustomVerificationEmailCmd.MarkFlagRequired("email-address")
		ses_sendCustomVerificationEmailCmd.MarkFlagRequired("template-name")
	})
	sesCmd.AddCommand(ses_sendCustomVerificationEmailCmd)
}

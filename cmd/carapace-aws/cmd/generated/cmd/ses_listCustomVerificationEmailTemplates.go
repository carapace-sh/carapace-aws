package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_listCustomVerificationEmailTemplatesCmd = &cobra.Command{
	Use:   "list-custom-verification-email-templates",
	Short: "Lists the existing custom verification email templates for your account in the current Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_listCustomVerificationEmailTemplatesCmd).Standalone()

	ses_listCustomVerificationEmailTemplatesCmd.Flags().String("max-results", "", "The maximum number of custom verification email templates to return.")
	ses_listCustomVerificationEmailTemplatesCmd.Flags().String("next-token", "", "An array the contains the name and creation time stamp for each template in your Amazon SES account.")
	sesCmd.AddCommand(ses_listCustomVerificationEmailTemplatesCmd)
}

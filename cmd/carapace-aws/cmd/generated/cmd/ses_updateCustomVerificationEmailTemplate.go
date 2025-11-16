package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_updateCustomVerificationEmailTemplateCmd = &cobra.Command{
	Use:   "update-custom-verification-email-template",
	Short: "Updates an existing custom verification email template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_updateCustomVerificationEmailTemplateCmd).Standalone()

	ses_updateCustomVerificationEmailTemplateCmd.Flags().String("failure-redirection-url", "", "The URL that the recipient of the verification email is sent to if his or her address is not successfully verified.")
	ses_updateCustomVerificationEmailTemplateCmd.Flags().String("from-email-address", "", "The email address that the custom verification email is sent from.")
	ses_updateCustomVerificationEmailTemplateCmd.Flags().String("success-redirection-url", "", "The URL that the recipient of the verification email is sent to if his or her address is successfully verified.")
	ses_updateCustomVerificationEmailTemplateCmd.Flags().String("template-content", "", "The content of the custom verification email.")
	ses_updateCustomVerificationEmailTemplateCmd.Flags().String("template-name", "", "The name of the custom verification email template to update.")
	ses_updateCustomVerificationEmailTemplateCmd.Flags().String("template-subject", "", "The subject line of the custom verification email.")
	ses_updateCustomVerificationEmailTemplateCmd.MarkFlagRequired("template-name")
	sesCmd.AddCommand(ses_updateCustomVerificationEmailTemplateCmd)
}

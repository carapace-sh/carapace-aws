package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_createCustomVerificationEmailTemplateCmd = &cobra.Command{
	Use:   "create-custom-verification-email-template",
	Short: "Creates a new custom verification email template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_createCustomVerificationEmailTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_createCustomVerificationEmailTemplateCmd).Standalone()

		ses_createCustomVerificationEmailTemplateCmd.Flags().String("failure-redirection-url", "", "The URL that the recipient of the verification email is sent to if his or her address is not successfully verified.")
		ses_createCustomVerificationEmailTemplateCmd.Flags().String("from-email-address", "", "The email address that the custom verification email is sent from.")
		ses_createCustomVerificationEmailTemplateCmd.Flags().String("success-redirection-url", "", "The URL that the recipient of the verification email is sent to if his or her address is successfully verified.")
		ses_createCustomVerificationEmailTemplateCmd.Flags().String("template-content", "", "The content of the custom verification email.")
		ses_createCustomVerificationEmailTemplateCmd.Flags().String("template-name", "", "The name of the custom verification email template.")
		ses_createCustomVerificationEmailTemplateCmd.Flags().String("template-subject", "", "The subject line of the custom verification email.")
		ses_createCustomVerificationEmailTemplateCmd.MarkFlagRequired("failure-redirection-url")
		ses_createCustomVerificationEmailTemplateCmd.MarkFlagRequired("from-email-address")
		ses_createCustomVerificationEmailTemplateCmd.MarkFlagRequired("success-redirection-url")
		ses_createCustomVerificationEmailTemplateCmd.MarkFlagRequired("template-content")
		ses_createCustomVerificationEmailTemplateCmd.MarkFlagRequired("template-name")
		ses_createCustomVerificationEmailTemplateCmd.MarkFlagRequired("template-subject")
	})
	sesCmd.AddCommand(ses_createCustomVerificationEmailTemplateCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_updateCustomVerificationEmailTemplateCmd = &cobra.Command{
	Use:   "update-custom-verification-email-template",
	Short: "Updates an existing custom verification email template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_updateCustomVerificationEmailTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_updateCustomVerificationEmailTemplateCmd).Standalone()

		sesv2_updateCustomVerificationEmailTemplateCmd.Flags().String("failure-redirection-url", "", "The URL that the recipient of the verification email is sent to if his or her address is not successfully verified.")
		sesv2_updateCustomVerificationEmailTemplateCmd.Flags().String("from-email-address", "", "The email address that the custom verification email is sent from.")
		sesv2_updateCustomVerificationEmailTemplateCmd.Flags().String("success-redirection-url", "", "The URL that the recipient of the verification email is sent to if his or her address is successfully verified.")
		sesv2_updateCustomVerificationEmailTemplateCmd.Flags().String("template-content", "", "The content of the custom verification email.")
		sesv2_updateCustomVerificationEmailTemplateCmd.Flags().String("template-name", "", "The name of the custom verification email template that you want to update.")
		sesv2_updateCustomVerificationEmailTemplateCmd.Flags().String("template-subject", "", "The subject line of the custom verification email.")
		sesv2_updateCustomVerificationEmailTemplateCmd.MarkFlagRequired("failure-redirection-url")
		sesv2_updateCustomVerificationEmailTemplateCmd.MarkFlagRequired("from-email-address")
		sesv2_updateCustomVerificationEmailTemplateCmd.MarkFlagRequired("success-redirection-url")
		sesv2_updateCustomVerificationEmailTemplateCmd.MarkFlagRequired("template-content")
		sesv2_updateCustomVerificationEmailTemplateCmd.MarkFlagRequired("template-name")
		sesv2_updateCustomVerificationEmailTemplateCmd.MarkFlagRequired("template-subject")
	})
	sesv2Cmd.AddCommand(sesv2_updateCustomVerificationEmailTemplateCmd)
}

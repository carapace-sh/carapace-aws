package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_createCustomVerificationEmailTemplateCmd = &cobra.Command{
	Use:   "create-custom-verification-email-template",
	Short: "Creates a new custom verification email template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_createCustomVerificationEmailTemplateCmd).Standalone()

	sesv2_createCustomVerificationEmailTemplateCmd.Flags().String("failure-redirection-url", "", "The URL that the recipient of the verification email is sent to if his or her address is not successfully verified.")
	sesv2_createCustomVerificationEmailTemplateCmd.Flags().String("from-email-address", "", "The email address that the custom verification email is sent from.")
	sesv2_createCustomVerificationEmailTemplateCmd.Flags().String("success-redirection-url", "", "The URL that the recipient of the verification email is sent to if his or her address is successfully verified.")
	sesv2_createCustomVerificationEmailTemplateCmd.Flags().String("template-content", "", "The content of the custom verification email.")
	sesv2_createCustomVerificationEmailTemplateCmd.Flags().String("template-name", "", "The name of the custom verification email template.")
	sesv2_createCustomVerificationEmailTemplateCmd.Flags().String("template-subject", "", "The subject line of the custom verification email.")
	sesv2_createCustomVerificationEmailTemplateCmd.MarkFlagRequired("failure-redirection-url")
	sesv2_createCustomVerificationEmailTemplateCmd.MarkFlagRequired("from-email-address")
	sesv2_createCustomVerificationEmailTemplateCmd.MarkFlagRequired("success-redirection-url")
	sesv2_createCustomVerificationEmailTemplateCmd.MarkFlagRequired("template-content")
	sesv2_createCustomVerificationEmailTemplateCmd.MarkFlagRequired("template-name")
	sesv2_createCustomVerificationEmailTemplateCmd.MarkFlagRequired("template-subject")
	sesv2Cmd.AddCommand(sesv2_createCustomVerificationEmailTemplateCmd)
}

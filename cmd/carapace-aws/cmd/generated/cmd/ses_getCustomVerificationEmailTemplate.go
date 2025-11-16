package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_getCustomVerificationEmailTemplateCmd = &cobra.Command{
	Use:   "get-custom-verification-email-template",
	Short: "Returns the custom email verification template for the template name you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_getCustomVerificationEmailTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_getCustomVerificationEmailTemplateCmd).Standalone()

		ses_getCustomVerificationEmailTemplateCmd.Flags().String("template-name", "", "The name of the custom verification email template to retrieve.")
		ses_getCustomVerificationEmailTemplateCmd.MarkFlagRequired("template-name")
	})
	sesCmd.AddCommand(ses_getCustomVerificationEmailTemplateCmd)
}

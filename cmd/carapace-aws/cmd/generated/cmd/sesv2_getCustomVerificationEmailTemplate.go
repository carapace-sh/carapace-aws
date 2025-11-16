package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getCustomVerificationEmailTemplateCmd = &cobra.Command{
	Use:   "get-custom-verification-email-template",
	Short: "Returns the custom email verification template for the template name you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getCustomVerificationEmailTemplateCmd).Standalone()

	sesv2_getCustomVerificationEmailTemplateCmd.Flags().String("template-name", "", "The name of the custom verification email template that you want to retrieve.")
	sesv2_getCustomVerificationEmailTemplateCmd.MarkFlagRequired("template-name")
	sesv2Cmd.AddCommand(sesv2_getCustomVerificationEmailTemplateCmd)
}

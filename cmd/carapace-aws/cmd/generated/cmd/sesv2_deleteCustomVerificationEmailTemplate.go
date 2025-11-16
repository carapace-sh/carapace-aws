package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_deleteCustomVerificationEmailTemplateCmd = &cobra.Command{
	Use:   "delete-custom-verification-email-template",
	Short: "Deletes an existing custom verification email template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_deleteCustomVerificationEmailTemplateCmd).Standalone()

	sesv2_deleteCustomVerificationEmailTemplateCmd.Flags().String("template-name", "", "The name of the custom verification email template that you want to delete.")
	sesv2_deleteCustomVerificationEmailTemplateCmd.MarkFlagRequired("template-name")
	sesv2Cmd.AddCommand(sesv2_deleteCustomVerificationEmailTemplateCmd)
}

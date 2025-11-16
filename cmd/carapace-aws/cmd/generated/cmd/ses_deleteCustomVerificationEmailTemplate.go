package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_deleteCustomVerificationEmailTemplateCmd = &cobra.Command{
	Use:   "delete-custom-verification-email-template",
	Short: "Deletes an existing custom verification email template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_deleteCustomVerificationEmailTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_deleteCustomVerificationEmailTemplateCmd).Standalone()

		ses_deleteCustomVerificationEmailTemplateCmd.Flags().String("template-name", "", "The name of the custom verification email template to delete.")
		ses_deleteCustomVerificationEmailTemplateCmd.MarkFlagRequired("template-name")
	})
	sesCmd.AddCommand(ses_deleteCustomVerificationEmailTemplateCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_createEmailTemplateCmd = &cobra.Command{
	Use:   "create-email-template",
	Short: "Creates an email template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_createEmailTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_createEmailTemplateCmd).Standalone()

		sesv2_createEmailTemplateCmd.Flags().String("template-content", "", "The content of the email template, composed of a subject line, an HTML part, and a text-only part.")
		sesv2_createEmailTemplateCmd.Flags().String("template-name", "", "The name of the template.")
		sesv2_createEmailTemplateCmd.MarkFlagRequired("template-content")
		sesv2_createEmailTemplateCmd.MarkFlagRequired("template-name")
	})
	sesv2Cmd.AddCommand(sesv2_createEmailTemplateCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_updateEmailTemplateCmd = &cobra.Command{
	Use:   "update-email-template",
	Short: "Updates an email template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_updateEmailTemplateCmd).Standalone()

	sesv2_updateEmailTemplateCmd.Flags().String("template-content", "", "The content of the email template, composed of a subject line, an HTML part, and a text-only part.")
	sesv2_updateEmailTemplateCmd.Flags().String("template-name", "", "The name of the template.")
	sesv2_updateEmailTemplateCmd.MarkFlagRequired("template-content")
	sesv2_updateEmailTemplateCmd.MarkFlagRequired("template-name")
	sesv2Cmd.AddCommand(sesv2_updateEmailTemplateCmd)
}

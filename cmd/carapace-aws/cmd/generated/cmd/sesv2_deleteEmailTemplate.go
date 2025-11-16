package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_deleteEmailTemplateCmd = &cobra.Command{
	Use:   "delete-email-template",
	Short: "Deletes an email template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_deleteEmailTemplateCmd).Standalone()

	sesv2_deleteEmailTemplateCmd.Flags().String("template-name", "", "The name of the template to be deleted.")
	sesv2_deleteEmailTemplateCmd.MarkFlagRequired("template-name")
	sesv2Cmd.AddCommand(sesv2_deleteEmailTemplateCmd)
}

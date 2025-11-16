package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getEmailTemplateCmd = &cobra.Command{
	Use:   "get-email-template",
	Short: "Displays the template object (which includes the subject line, HTML part and text part) for the template you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getEmailTemplateCmd).Standalone()

	sesv2_getEmailTemplateCmd.Flags().String("template-name", "", "The name of the template.")
	sesv2_getEmailTemplateCmd.MarkFlagRequired("template-name")
	sesv2Cmd.AddCommand(sesv2_getEmailTemplateCmd)
}

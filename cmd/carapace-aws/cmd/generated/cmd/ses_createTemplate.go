package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_createTemplateCmd = &cobra.Command{
	Use:   "create-template",
	Short: "Creates an email template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_createTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_createTemplateCmd).Standalone()

		ses_createTemplateCmd.Flags().String("template", "", "The content of the email, composed of a subject line and either an HTML part or a text-only part.")
		ses_createTemplateCmd.MarkFlagRequired("template")
	})
	sesCmd.AddCommand(ses_createTemplateCmd)
}

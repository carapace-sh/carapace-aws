package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_deleteTemplateCmd = &cobra.Command{
	Use:   "delete-template",
	Short: "Deletes an email template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_deleteTemplateCmd).Standalone()

	ses_deleteTemplateCmd.Flags().String("template-name", "", "The name of the template to be deleted.")
	ses_deleteTemplateCmd.MarkFlagRequired("template-name")
	sesCmd.AddCommand(ses_deleteTemplateCmd)
}

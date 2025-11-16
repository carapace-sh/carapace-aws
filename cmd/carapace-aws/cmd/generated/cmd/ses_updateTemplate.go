package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_updateTemplateCmd = &cobra.Command{
	Use:   "update-template",
	Short: "Updates an email template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_updateTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_updateTemplateCmd).Standalone()

		ses_updateTemplateCmd.Flags().String("template", "", "")
		ses_updateTemplateCmd.MarkFlagRequired("template")
	})
	sesCmd.AddCommand(ses_updateTemplateCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_getTemplateCmd = &cobra.Command{
	Use:   "get-template",
	Short: "Displays the template object (which includes the Subject line, HTML part and text part) for the template you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_getTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_getTemplateCmd).Standalone()

		ses_getTemplateCmd.Flags().String("template-name", "", "The name of the template to retrieve.")
		ses_getTemplateCmd.MarkFlagRequired("template-name")
	})
	sesCmd.AddCommand(ses_getTemplateCmd)
}

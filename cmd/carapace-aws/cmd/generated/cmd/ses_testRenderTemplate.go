package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_testRenderTemplateCmd = &cobra.Command{
	Use:   "test-render-template",
	Short: "Creates a preview of the MIME content of an email when provided with a template and a set of replacement data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_testRenderTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_testRenderTemplateCmd).Standalone()

		ses_testRenderTemplateCmd.Flags().String("template-data", "", "A list of replacement values to apply to the template.")
		ses_testRenderTemplateCmd.Flags().String("template-name", "", "The name of the template to render.")
		ses_testRenderTemplateCmd.MarkFlagRequired("template-data")
		ses_testRenderTemplateCmd.MarkFlagRequired("template-name")
	})
	sesCmd.AddCommand(ses_testRenderTemplateCmd)
}

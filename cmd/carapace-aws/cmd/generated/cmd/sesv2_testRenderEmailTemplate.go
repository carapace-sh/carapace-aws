package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_testRenderEmailTemplateCmd = &cobra.Command{
	Use:   "test-render-email-template",
	Short: "Creates a preview of the MIME content of an email when provided with a template and a set of replacement data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_testRenderEmailTemplateCmd).Standalone()

	sesv2_testRenderEmailTemplateCmd.Flags().String("template-data", "", "A list of replacement values to apply to the template.")
	sesv2_testRenderEmailTemplateCmd.Flags().String("template-name", "", "The name of the template.")
	sesv2_testRenderEmailTemplateCmd.MarkFlagRequired("template-data")
	sesv2_testRenderEmailTemplateCmd.MarkFlagRequired("template-name")
	sesv2Cmd.AddCommand(sesv2_testRenderEmailTemplateCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_getGeneratedTemplateCmd = &cobra.Command{
	Use:   "get-generated-template",
	Short: "Retrieves a generated template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_getGeneratedTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_getGeneratedTemplateCmd).Standalone()

		cloudformation_getGeneratedTemplateCmd.Flags().String("format", "", "The language to use to retrieve for the generated template.")
		cloudformation_getGeneratedTemplateCmd.Flags().String("generated-template-name", "", "The name or Amazon Resource Name (ARN) of the generated template.")
		cloudformation_getGeneratedTemplateCmd.MarkFlagRequired("generated-template-name")
	})
	cloudformationCmd.AddCommand(cloudformation_getGeneratedTemplateCmd)
}

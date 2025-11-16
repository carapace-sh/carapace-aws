package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_describeGeneratedTemplateCmd = &cobra.Command{
	Use:   "describe-generated-template",
	Short: "Describes a generated template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_describeGeneratedTemplateCmd).Standalone()

	cloudformation_describeGeneratedTemplateCmd.Flags().String("generated-template-name", "", "The name or Amazon Resource Name (ARN) of a generated template.")
	cloudformation_describeGeneratedTemplateCmd.MarkFlagRequired("generated-template-name")
	cloudformationCmd.AddCommand(cloudformation_describeGeneratedTemplateCmd)
}

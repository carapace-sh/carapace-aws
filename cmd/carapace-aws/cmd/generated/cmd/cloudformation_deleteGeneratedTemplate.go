package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_deleteGeneratedTemplateCmd = &cobra.Command{
	Use:   "delete-generated-template",
	Short: "Deleted a generated template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_deleteGeneratedTemplateCmd).Standalone()

	cloudformation_deleteGeneratedTemplateCmd.Flags().String("generated-template-name", "", "The name or Amazon Resource Name (ARN) of a generated template.")
	cloudformation_deleteGeneratedTemplateCmd.MarkFlagRequired("generated-template-name")
	cloudformationCmd.AddCommand(cloudformation_deleteGeneratedTemplateCmd)
}

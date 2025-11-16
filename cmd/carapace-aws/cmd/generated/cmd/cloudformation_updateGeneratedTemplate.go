package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_updateGeneratedTemplateCmd = &cobra.Command{
	Use:   "update-generated-template",
	Short: "Updates a generated template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_updateGeneratedTemplateCmd).Standalone()

	cloudformation_updateGeneratedTemplateCmd.Flags().String("add-resources", "", "An optional list of resources to be added to the generated template.")
	cloudformation_updateGeneratedTemplateCmd.Flags().String("generated-template-name", "", "The name or Amazon Resource Name (ARN) of a generated template.")
	cloudformation_updateGeneratedTemplateCmd.Flags().String("new-generated-template-name", "", "An optional new name to assign to the generated template.")
	cloudformation_updateGeneratedTemplateCmd.Flags().String("refresh-all-resources", "", "If `true`, update the resource properties in the generated template with their current live state.")
	cloudformation_updateGeneratedTemplateCmd.Flags().String("remove-resources", "", "A list of logical ids for resources to remove from the generated template.")
	cloudformation_updateGeneratedTemplateCmd.Flags().String("template-configuration", "", "The configuration details of the generated template, including the `DeletionPolicy` and `UpdateReplacePolicy`.")
	cloudformation_updateGeneratedTemplateCmd.MarkFlagRequired("generated-template-name")
	cloudformationCmd.AddCommand(cloudformation_updateGeneratedTemplateCmd)
}

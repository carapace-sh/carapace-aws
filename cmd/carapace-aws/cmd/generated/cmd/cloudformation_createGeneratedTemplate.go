package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_createGeneratedTemplateCmd = &cobra.Command{
	Use:   "create-generated-template",
	Short: "Creates a template from existing resources that are not already managed with CloudFormation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_createGeneratedTemplateCmd).Standalone()

	cloudformation_createGeneratedTemplateCmd.Flags().String("generated-template-name", "", "The name assigned to the generated template.")
	cloudformation_createGeneratedTemplateCmd.Flags().String("resources", "", "An optional list of resources to be included in the generated template.")
	cloudformation_createGeneratedTemplateCmd.Flags().String("stack-name", "", "An optional name or ARN of a stack to use as the base stack for the generated template.")
	cloudformation_createGeneratedTemplateCmd.Flags().String("template-configuration", "", "The configuration details of the generated template, including the `DeletionPolicy` and `UpdateReplacePolicy`.")
	cloudformation_createGeneratedTemplateCmd.MarkFlagRequired("generated-template-name")
	cloudformationCmd.AddCommand(cloudformation_createGeneratedTemplateCmd)
}

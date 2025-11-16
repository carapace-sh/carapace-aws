package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_getTemplateCmd = &cobra.Command{
	Use:   "get-template",
	Short: "Returns the template body for a specified stack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_getTemplateCmd).Standalone()

	cloudformation_getTemplateCmd.Flags().String("change-set-name", "", "The name or Amazon Resource Name (ARN) of a change set for which CloudFormation returns the associated template.")
	cloudformation_getTemplateCmd.Flags().String("stack-name", "", "The name or the unique stack ID that's associated with the stack, which aren't always interchangeable:")
	cloudformation_getTemplateCmd.Flags().String("template-stage", "", "For templates that include transforms, the stage of the template that CloudFormation returns.")
	cloudformationCmd.AddCommand(cloudformation_getTemplateCmd)
}

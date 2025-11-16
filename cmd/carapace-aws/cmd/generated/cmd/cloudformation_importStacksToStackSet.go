package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_importStacksToStackSetCmd = &cobra.Command{
	Use:   "import-stacks-to-stack-set",
	Short: "Import existing stacks into a new StackSets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_importStacksToStackSetCmd).Standalone()

	cloudformation_importStacksToStackSetCmd.Flags().String("call-as", "", "By default, `SELF` is specified.")
	cloudformation_importStacksToStackSetCmd.Flags().String("operation-id", "", "A unique, user defined, identifier for the StackSet operation.")
	cloudformation_importStacksToStackSetCmd.Flags().String("operation-preferences", "", "The user-specified preferences for how CloudFormation performs a StackSet operation.")
	cloudformation_importStacksToStackSetCmd.Flags().String("organizational-unit-ids", "", "The list of OU ID's to which the imported stacks must be mapped as deployment targets.")
	cloudformation_importStacksToStackSetCmd.Flags().String("stack-ids", "", "The IDs of the stacks you are importing into a StackSet.")
	cloudformation_importStacksToStackSetCmd.Flags().String("stack-ids-url", "", "The Amazon S3 URL which contains list of stack ids to be inputted.")
	cloudformation_importStacksToStackSetCmd.Flags().String("stack-set-name", "", "The name of the StackSet.")
	cloudformation_importStacksToStackSetCmd.MarkFlagRequired("stack-set-name")
	cloudformationCmd.AddCommand(cloudformation_importStacksToStackSetCmd)
}

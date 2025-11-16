package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_createStackRefactorCmd = &cobra.Command{
	Use:   "create-stack-refactor",
	Short: "Creates a refactor across multiple stacks, with the list of stacks and resources that are affected.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_createStackRefactorCmd).Standalone()

	cloudformation_createStackRefactorCmd.Flags().String("description", "", "A description to help you identify the stack refactor.")
	cloudformation_createStackRefactorCmd.Flags().String("enable-stack-creation", "", "Determines if a new stack is created with the refactor.")
	cloudformation_createStackRefactorCmd.Flags().String("resource-mappings", "", "The mappings for the stack resource `Source` and stack resource `Destination`.")
	cloudformation_createStackRefactorCmd.Flags().String("stack-definitions", "", "The stacks being refactored.")
	cloudformation_createStackRefactorCmd.MarkFlagRequired("stack-definitions")
	cloudformationCmd.AddCommand(cloudformation_createStackRefactorCmd)
}

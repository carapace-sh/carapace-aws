package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_executeStackRefactorCmd = &cobra.Command{
	Use:   "execute-stack-refactor",
	Short: "Executes the stack refactor operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_executeStackRefactorCmd).Standalone()

	cloudformation_executeStackRefactorCmd.Flags().String("stack-refactor-id", "", "The ID associated with the stack refactor created from the [CreateStackRefactor]() action.")
	cloudformation_executeStackRefactorCmd.MarkFlagRequired("stack-refactor-id")
	cloudformationCmd.AddCommand(cloudformation_executeStackRefactorCmd)
}

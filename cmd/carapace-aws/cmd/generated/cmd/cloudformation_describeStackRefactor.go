package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_describeStackRefactorCmd = &cobra.Command{
	Use:   "describe-stack-refactor",
	Short: "Describes the stack refactor status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_describeStackRefactorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_describeStackRefactorCmd).Standalone()

		cloudformation_describeStackRefactorCmd.Flags().String("stack-refactor-id", "", "The ID associated with the stack refactor created from the [CreateStackRefactor]() action.")
		cloudformation_describeStackRefactorCmd.MarkFlagRequired("stack-refactor-id")
	})
	cloudformationCmd.AddCommand(cloudformation_describeStackRefactorCmd)
}

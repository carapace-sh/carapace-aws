package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_deleteEvaluationCmd = &cobra.Command{
	Use:   "delete-evaluation",
	Short: "Assigns the `DELETED` status to an `Evaluation`, rendering it unusable.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_deleteEvaluationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(machinelearning_deleteEvaluationCmd).Standalone()

		machinelearning_deleteEvaluationCmd.Flags().String("evaluation-id", "", "A user-supplied ID that uniquely identifies the `Evaluation` to delete.")
		machinelearning_deleteEvaluationCmd.MarkFlagRequired("evaluation-id")
	})
	machinelearningCmd.AddCommand(machinelearning_deleteEvaluationCmd)
}

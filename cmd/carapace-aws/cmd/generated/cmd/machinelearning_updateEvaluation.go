package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_updateEvaluationCmd = &cobra.Command{
	Use:   "update-evaluation",
	Short: "Updates the `EvaluationName` of an `Evaluation`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_updateEvaluationCmd).Standalone()

	machinelearning_updateEvaluationCmd.Flags().String("evaluation-id", "", "The ID assigned to the `Evaluation` during creation.")
	machinelearning_updateEvaluationCmd.Flags().String("evaluation-name", "", "A new user-supplied name or description of the `Evaluation` that will replace the current content.")
	machinelearning_updateEvaluationCmd.MarkFlagRequired("evaluation-id")
	machinelearning_updateEvaluationCmd.MarkFlagRequired("evaluation-name")
	machinelearningCmd.AddCommand(machinelearning_updateEvaluationCmd)
}

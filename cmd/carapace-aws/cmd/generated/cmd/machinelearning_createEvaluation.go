package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_createEvaluationCmd = &cobra.Command{
	Use:   "create-evaluation",
	Short: "Creates a new `Evaluation` of an `MLModel`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_createEvaluationCmd).Standalone()

	machinelearning_createEvaluationCmd.Flags().String("evaluation-data-source-id", "", "The ID of the `DataSource` for the evaluation.")
	machinelearning_createEvaluationCmd.Flags().String("evaluation-id", "", "A user-supplied ID that uniquely identifies the `Evaluation`.")
	machinelearning_createEvaluationCmd.Flags().String("evaluation-name", "", "A user-supplied name or description of the `Evaluation`.")
	machinelearning_createEvaluationCmd.Flags().String("mlmodel-id", "", "The ID of the `MLModel` to evaluate.")
	machinelearning_createEvaluationCmd.MarkFlagRequired("evaluation-data-source-id")
	machinelearning_createEvaluationCmd.MarkFlagRequired("evaluation-id")
	machinelearning_createEvaluationCmd.MarkFlagRequired("mlmodel-id")
	machinelearningCmd.AddCommand(machinelearning_createEvaluationCmd)
}

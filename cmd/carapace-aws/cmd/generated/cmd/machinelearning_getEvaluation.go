package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_getEvaluationCmd = &cobra.Command{
	Use:   "get-evaluation",
	Short: "Returns an `Evaluation` that includes metadata as well as the current status of the `Evaluation`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_getEvaluationCmd).Standalone()

	machinelearning_getEvaluationCmd.Flags().String("evaluation-id", "", "The ID of the `Evaluation` to retrieve.")
	machinelearning_getEvaluationCmd.MarkFlagRequired("evaluation-id")
	machinelearningCmd.AddCommand(machinelearning_getEvaluationCmd)
}

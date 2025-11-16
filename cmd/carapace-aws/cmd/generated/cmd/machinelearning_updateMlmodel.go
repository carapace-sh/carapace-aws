package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_updateMlmodelCmd = &cobra.Command{
	Use:   "update-mlmodel",
	Short: "Updates the `MLModelName` and the `ScoreThreshold` of an `MLModel`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_updateMlmodelCmd).Standalone()

	machinelearning_updateMlmodelCmd.Flags().String("mlmodel-id", "", "The ID assigned to the `MLModel` during creation.")
	machinelearning_updateMlmodelCmd.Flags().String("mlmodel-name", "", "A user-supplied name or description of the `MLModel`.")
	machinelearning_updateMlmodelCmd.Flags().String("score-threshold", "", "The `ScoreThreshold` used in binary classification `MLModel` that marks the boundary between a positive prediction and a negative prediction.")
	machinelearning_updateMlmodelCmd.MarkFlagRequired("mlmodel-id")
	machinelearningCmd.AddCommand(machinelearning_updateMlmodelCmd)
}

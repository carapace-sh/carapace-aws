package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_predictCmd = &cobra.Command{
	Use:   "predict",
	Short: "Generates a prediction for the observation using the specified `ML Model`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_predictCmd).Standalone()

	machinelearning_predictCmd.Flags().String("mlmodel-id", "", "A unique identifier of the `MLModel`.")
	machinelearning_predictCmd.Flags().String("predict-endpoint", "", "")
	machinelearning_predictCmd.Flags().String("record", "", "")
	machinelearning_predictCmd.MarkFlagRequired("mlmodel-id")
	machinelearning_predictCmd.MarkFlagRequired("predict-endpoint")
	machinelearning_predictCmd.MarkFlagRequired("record")
	machinelearningCmd.AddCommand(machinelearning_predictCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_createMlmodelCmd = &cobra.Command{
	Use:   "create-mlmodel",
	Short: "Creates a new `MLModel` using the `DataSource` and the recipe as information sources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_createMlmodelCmd).Standalone()

	machinelearning_createMlmodelCmd.Flags().String("mlmodel-id", "", "A user-supplied ID that uniquely identifies the `MLModel`.")
	machinelearning_createMlmodelCmd.Flags().String("mlmodel-name", "", "A user-supplied name or description of the `MLModel`.")
	machinelearning_createMlmodelCmd.Flags().String("mlmodel-type", "", "The category of supervised learning that this `MLModel` will address.")
	machinelearning_createMlmodelCmd.Flags().String("parameters", "", "A list of the training parameters in the `MLModel`.")
	machinelearning_createMlmodelCmd.Flags().String("recipe", "", "The data recipe for creating the `MLModel`.")
	machinelearning_createMlmodelCmd.Flags().String("recipe-uri", "", "The Amazon Simple Storage Service (Amazon S3) location and file name that contains the `MLModel` recipe.")
	machinelearning_createMlmodelCmd.Flags().String("training-data-source-id", "", "The `DataSource` that points to the training data.")
	machinelearning_createMlmodelCmd.MarkFlagRequired("mlmodel-id")
	machinelearning_createMlmodelCmd.MarkFlagRequired("mlmodel-type")
	machinelearning_createMlmodelCmd.MarkFlagRequired("training-data-source-id")
	machinelearningCmd.AddCommand(machinelearning_createMlmodelCmd)
}

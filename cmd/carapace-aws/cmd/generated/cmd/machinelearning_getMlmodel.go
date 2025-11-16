package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_getMlmodelCmd = &cobra.Command{
	Use:   "get-mlmodel",
	Short: "Returns an `MLModel` that includes detailed metadata, data source information, and the current status of the `MLModel`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_getMlmodelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(machinelearning_getMlmodelCmd).Standalone()

		machinelearning_getMlmodelCmd.Flags().String("mlmodel-id", "", "The ID assigned to the `MLModel` at creation.")
		machinelearning_getMlmodelCmd.Flags().String("verbose", "", "Specifies whether the `GetMLModel` operation should return `Recipe`.")
		machinelearning_getMlmodelCmd.MarkFlagRequired("mlmodel-id")
	})
	machinelearningCmd.AddCommand(machinelearning_getMlmodelCmd)
}

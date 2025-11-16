package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_createComputationModelCmd = &cobra.Command{
	Use:   "create-computation-model",
	Short: "Create a computation model with a configuration and data binding.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_createComputationModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_createComputationModelCmd).Standalone()

		iotsitewise_createComputationModelCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iotsitewise_createComputationModelCmd.Flags().String("computation-model-configuration", "", "The configuration for the computation model.")
		iotsitewise_createComputationModelCmd.Flags().String("computation-model-data-binding", "", "The data binding for the computation model.")
		iotsitewise_createComputationModelCmd.Flags().String("computation-model-description", "", "The description of the computation model.")
		iotsitewise_createComputationModelCmd.Flags().String("computation-model-name", "", "The name of the computation model.")
		iotsitewise_createComputationModelCmd.Flags().String("tags", "", "A list of key-value pairs that contain metadata for the asset.")
		iotsitewise_createComputationModelCmd.MarkFlagRequired("computation-model-configuration")
		iotsitewise_createComputationModelCmd.MarkFlagRequired("computation-model-data-binding")
		iotsitewise_createComputationModelCmd.MarkFlagRequired("computation-model-name")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_createComputationModelCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_updateComputationModelCmd = &cobra.Command{
	Use:   "update-computation-model",
	Short: "Updates the computation model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_updateComputationModelCmd).Standalone()

	iotsitewise_updateComputationModelCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iotsitewise_updateComputationModelCmd.Flags().String("computation-model-configuration", "", "The configuration for the computation model.")
	iotsitewise_updateComputationModelCmd.Flags().String("computation-model-data-binding", "", "The data binding for the computation model.")
	iotsitewise_updateComputationModelCmd.Flags().String("computation-model-description", "", "The description of the computation model.")
	iotsitewise_updateComputationModelCmd.Flags().String("computation-model-id", "", "The ID of the computation model.")
	iotsitewise_updateComputationModelCmd.Flags().String("computation-model-name", "", "The name of the computation model.")
	iotsitewise_updateComputationModelCmd.MarkFlagRequired("computation-model-configuration")
	iotsitewise_updateComputationModelCmd.MarkFlagRequired("computation-model-data-binding")
	iotsitewise_updateComputationModelCmd.MarkFlagRequired("computation-model-id")
	iotsitewise_updateComputationModelCmd.MarkFlagRequired("computation-model-name")
	iotsitewiseCmd.AddCommand(iotsitewise_updateComputationModelCmd)
}

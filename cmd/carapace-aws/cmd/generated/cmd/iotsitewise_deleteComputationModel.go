package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_deleteComputationModelCmd = &cobra.Command{
	Use:   "delete-computation-model",
	Short: "Deletes a computation model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_deleteComputationModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_deleteComputationModelCmd).Standalone()

		iotsitewise_deleteComputationModelCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iotsitewise_deleteComputationModelCmd.Flags().String("computation-model-id", "", "The ID of the computation model.")
		iotsitewise_deleteComputationModelCmd.MarkFlagRequired("computation-model-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_deleteComputationModelCmd)
}

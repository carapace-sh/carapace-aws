package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describeComputationModelCmd = &cobra.Command{
	Use:   "describe-computation-model",
	Short: "Retrieves information about a computation model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describeComputationModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_describeComputationModelCmd).Standalone()

		iotsitewise_describeComputationModelCmd.Flags().String("computation-model-id", "", "The ID of the computation model.")
		iotsitewise_describeComputationModelCmd.Flags().String("computation-model-version", "", "The version of the computation model.")
		iotsitewise_describeComputationModelCmd.MarkFlagRequired("computation-model-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_describeComputationModelCmd)
}

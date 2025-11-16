package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describeComputationModelExecutionSummaryCmd = &cobra.Command{
	Use:   "describe-computation-model-execution-summary",
	Short: "Retrieves information about the execution summary of a computation model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describeComputationModelExecutionSummaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_describeComputationModelExecutionSummaryCmd).Standalone()

		iotsitewise_describeComputationModelExecutionSummaryCmd.Flags().String("computation-model-id", "", "The ID of the computation model.")
		iotsitewise_describeComputationModelExecutionSummaryCmd.Flags().String("resolve-to-resource-id", "", "The ID of the resolved resource.")
		iotsitewise_describeComputationModelExecutionSummaryCmd.Flags().String("resolve-to-resource-type", "", "The type of the resolved resource.")
		iotsitewise_describeComputationModelExecutionSummaryCmd.MarkFlagRequired("computation-model-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_describeComputationModelExecutionSummaryCmd)
}

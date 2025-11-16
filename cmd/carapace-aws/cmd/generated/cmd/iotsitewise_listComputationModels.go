package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listComputationModelsCmd = &cobra.Command{
	Use:   "list-computation-models",
	Short: "Retrieves a paginated list of summaries of all computation models.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listComputationModelsCmd).Standalone()

	iotsitewise_listComputationModelsCmd.Flags().String("computation-model-type", "", "The type of computation model.")
	iotsitewise_listComputationModelsCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
	iotsitewise_listComputationModelsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	iotsitewiseCmd.AddCommand(iotsitewise_listComputationModelsCmd)
}

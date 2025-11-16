package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_describeBatchPredictionsCmd = &cobra.Command{
	Use:   "describe-batch-predictions",
	Short: "Returns a list of `BatchPrediction` operations that match the search criteria in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_describeBatchPredictionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(machinelearning_describeBatchPredictionsCmd).Standalone()

		machinelearning_describeBatchPredictionsCmd.Flags().String("eq", "", "The equal to operator.")
		machinelearning_describeBatchPredictionsCmd.Flags().String("filter-variable", "", "Use one of the following variables to filter a list of `BatchPrediction`:")
		machinelearning_describeBatchPredictionsCmd.Flags().String("ge", "", "The greater than or equal to operator.")
		machinelearning_describeBatchPredictionsCmd.Flags().String("gt", "", "The greater than operator.")
		machinelearning_describeBatchPredictionsCmd.Flags().String("le", "", "The less than or equal to operator.")
		machinelearning_describeBatchPredictionsCmd.Flags().String("limit", "", "The number of pages of information to include in the result.")
		machinelearning_describeBatchPredictionsCmd.Flags().String("lt", "", "The less than operator.")
		machinelearning_describeBatchPredictionsCmd.Flags().String("ne", "", "The not equal to operator.")
		machinelearning_describeBatchPredictionsCmd.Flags().String("next-token", "", "An ID of the page in the paginated results.")
		machinelearning_describeBatchPredictionsCmd.Flags().String("prefix", "", "A string that is found at the beginning of a variable, such as `Name` or `Id`.")
		machinelearning_describeBatchPredictionsCmd.Flags().String("sort-order", "", "A two-value parameter that determines the sequence of the resulting list of `MLModel`s.")
	})
	machinelearningCmd.AddCommand(machinelearning_describeBatchPredictionsCmd)
}

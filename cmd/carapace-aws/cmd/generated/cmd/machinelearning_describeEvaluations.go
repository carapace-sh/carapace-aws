package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_describeEvaluationsCmd = &cobra.Command{
	Use:   "describe-evaluations",
	Short: "Returns a list of `DescribeEvaluations` that match the search criteria in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_describeEvaluationsCmd).Standalone()

	machinelearning_describeEvaluationsCmd.Flags().String("eq", "", "The equal to operator.")
	machinelearning_describeEvaluationsCmd.Flags().String("filter-variable", "", "Use one of the following variable to filter a list of `Evaluation` objects:")
	machinelearning_describeEvaluationsCmd.Flags().String("ge", "", "The greater than or equal to operator.")
	machinelearning_describeEvaluationsCmd.Flags().String("gt", "", "The greater than operator.")
	machinelearning_describeEvaluationsCmd.Flags().String("le", "", "The less than or equal to operator.")
	machinelearning_describeEvaluationsCmd.Flags().String("limit", "", "The maximum number of `Evaluation` to include in the result.")
	machinelearning_describeEvaluationsCmd.Flags().String("lt", "", "The less than operator.")
	machinelearning_describeEvaluationsCmd.Flags().String("ne", "", "The not equal to operator.")
	machinelearning_describeEvaluationsCmd.Flags().String("next-token", "", "The ID of the page in the paginated results.")
	machinelearning_describeEvaluationsCmd.Flags().String("prefix", "", "A string that is found at the beginning of a variable, such as `Name` or `Id`.")
	machinelearning_describeEvaluationsCmd.Flags().String("sort-order", "", "A two-value parameter that determines the sequence of the resulting list of `Evaluation`.")
	machinelearningCmd.AddCommand(machinelearning_describeEvaluationsCmd)
}

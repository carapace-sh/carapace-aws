package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listAlgorithmsCmd = &cobra.Command{
	Use:   "list-algorithms",
	Short: "Lists the machine learning algorithms that have been created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listAlgorithmsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listAlgorithmsCmd).Standalone()

		sagemaker_listAlgorithmsCmd.Flags().String("creation-time-after", "", "A filter that returns only algorithms created after the specified time (timestamp).")
		sagemaker_listAlgorithmsCmd.Flags().String("creation-time-before", "", "A filter that returns only algorithms created before the specified time (timestamp).")
		sagemaker_listAlgorithmsCmd.Flags().String("max-results", "", "The maximum number of algorithms to return in the response.")
		sagemaker_listAlgorithmsCmd.Flags().String("name-contains", "", "A string in the algorithm name.")
		sagemaker_listAlgorithmsCmd.Flags().String("next-token", "", "If the response to a previous `ListAlgorithms` request was truncated, the response includes a `NextToken`.")
		sagemaker_listAlgorithmsCmd.Flags().String("sort-by", "", "The parameter by which to sort the results.")
		sagemaker_listAlgorithmsCmd.Flags().String("sort-order", "", "The sort order for the results.")
	})
	sagemakerCmd.AddCommand(sagemaker_listAlgorithmsCmd)
}

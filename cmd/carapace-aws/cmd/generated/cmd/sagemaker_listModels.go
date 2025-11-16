package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listModelsCmd = &cobra.Command{
	Use:   "list-models",
	Short: "Lists models created with the `CreateModel` API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listModelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listModelsCmd).Standalone()

		sagemaker_listModelsCmd.Flags().String("creation-time-after", "", "A filter that returns only models with a creation time greater than or equal to the specified time (timestamp).")
		sagemaker_listModelsCmd.Flags().String("creation-time-before", "", "A filter that returns only models created before the specified time (timestamp).")
		sagemaker_listModelsCmd.Flags().String("max-results", "", "The maximum number of models to return in the response.")
		sagemaker_listModelsCmd.Flags().String("name-contains", "", "A string in the model name.")
		sagemaker_listModelsCmd.Flags().String("next-token", "", "If the response to a previous `ListModels` request was truncated, the response includes a `NextToken`.")
		sagemaker_listModelsCmd.Flags().String("sort-by", "", "Sorts the list of results.")
		sagemaker_listModelsCmd.Flags().String("sort-order", "", "The sort order for results.")
	})
	sagemakerCmd.AddCommand(sagemaker_listModelsCmd)
}

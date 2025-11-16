package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listInferenceComponentsCmd = &cobra.Command{
	Use:   "list-inference-components",
	Short: "Lists the inference components in your account and their properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listInferenceComponentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listInferenceComponentsCmd).Standalone()

		sagemaker_listInferenceComponentsCmd.Flags().String("creation-time-after", "", "Filters the results to only those inference components that were created after the specified time.")
		sagemaker_listInferenceComponentsCmd.Flags().String("creation-time-before", "", "Filters the results to only those inference components that were created before the specified time.")
		sagemaker_listInferenceComponentsCmd.Flags().String("endpoint-name-equals", "", "An endpoint name to filter the listed inference components.")
		sagemaker_listInferenceComponentsCmd.Flags().String("last-modified-time-after", "", "Filters the results to only those inference components that were updated after the specified time.")
		sagemaker_listInferenceComponentsCmd.Flags().String("last-modified-time-before", "", "Filters the results to only those inference components that were updated before the specified time.")
		sagemaker_listInferenceComponentsCmd.Flags().String("max-results", "", "The maximum number of inference components to return in the response.")
		sagemaker_listInferenceComponentsCmd.Flags().String("name-contains", "", "Filters the results to only those inference components with a name that contains the specified string.")
		sagemaker_listInferenceComponentsCmd.Flags().String("next-token", "", "A token that you use to get the next set of results following a truncated response.")
		sagemaker_listInferenceComponentsCmd.Flags().String("sort-by", "", "The field by which to sort the inference components in the response.")
		sagemaker_listInferenceComponentsCmd.Flags().String("sort-order", "", "The sort order for results.")
		sagemaker_listInferenceComponentsCmd.Flags().String("status-equals", "", "Filters the results to only those inference components with the specified status.")
		sagemaker_listInferenceComponentsCmd.Flags().String("variant-name-equals", "", "A production variant name to filter the listed inference components.")
	})
	sagemakerCmd.AddCommand(sagemaker_listInferenceComponentsCmd)
}

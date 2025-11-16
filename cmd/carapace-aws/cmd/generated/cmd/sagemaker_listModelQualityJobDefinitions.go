package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listModelQualityJobDefinitionsCmd = &cobra.Command{
	Use:   "list-model-quality-job-definitions",
	Short: "Gets a list of model quality monitoring job definitions in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listModelQualityJobDefinitionsCmd).Standalone()

	sagemaker_listModelQualityJobDefinitionsCmd.Flags().String("creation-time-after", "", "A filter that returns only model quality monitoring job definitions created after the specified time.")
	sagemaker_listModelQualityJobDefinitionsCmd.Flags().String("creation-time-before", "", "A filter that returns only model quality monitoring job definitions created before the specified time.")
	sagemaker_listModelQualityJobDefinitionsCmd.Flags().String("endpoint-name", "", "A filter that returns only model quality monitoring job definitions that are associated with the specified endpoint.")
	sagemaker_listModelQualityJobDefinitionsCmd.Flags().String("max-results", "", "The maximum number of results to return in a call to `ListModelQualityJobDefinitions`.")
	sagemaker_listModelQualityJobDefinitionsCmd.Flags().String("name-contains", "", "A string in the transform job name.")
	sagemaker_listModelQualityJobDefinitionsCmd.Flags().String("next-token", "", "If the result of the previous `ListModelQualityJobDefinitions` request was truncated, the response includes a `NextToken`.")
	sagemaker_listModelQualityJobDefinitionsCmd.Flags().String("sort-by", "", "The field to sort results by.")
	sagemaker_listModelQualityJobDefinitionsCmd.Flags().String("sort-order", "", "Whether to sort the results in `Ascending` or `Descending` order.")
	sagemakerCmd.AddCommand(sagemaker_listModelQualityJobDefinitionsCmd)
}

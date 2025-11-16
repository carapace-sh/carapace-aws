package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listDataQualityJobDefinitionsCmd = &cobra.Command{
	Use:   "list-data-quality-job-definitions",
	Short: "Lists the data quality job definitions in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listDataQualityJobDefinitionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listDataQualityJobDefinitionsCmd).Standalone()

		sagemaker_listDataQualityJobDefinitionsCmd.Flags().String("creation-time-after", "", "A filter that returns only data quality monitoring job definitions created after the specified time.")
		sagemaker_listDataQualityJobDefinitionsCmd.Flags().String("creation-time-before", "", "A filter that returns only data quality monitoring job definitions created before the specified time.")
		sagemaker_listDataQualityJobDefinitionsCmd.Flags().String("endpoint-name", "", "A filter that lists the data quality job definitions associated with the specified endpoint.")
		sagemaker_listDataQualityJobDefinitionsCmd.Flags().String("max-results", "", "The maximum number of data quality monitoring job definitions to return in the response.")
		sagemaker_listDataQualityJobDefinitionsCmd.Flags().String("name-contains", "", "A string in the data quality monitoring job definition name.")
		sagemaker_listDataQualityJobDefinitionsCmd.Flags().String("next-token", "", "If the result of the previous `ListDataQualityJobDefinitions` request was truncated, the response includes a `NextToken`.")
		sagemaker_listDataQualityJobDefinitionsCmd.Flags().String("sort-by", "", "The field to sort results by.")
		sagemaker_listDataQualityJobDefinitionsCmd.Flags().String("sort-order", "", "Whether to sort the results in `Ascending` or `Descending` order.")
	})
	sagemakerCmd.AddCommand(sagemaker_listDataQualityJobDefinitionsCmd)
}

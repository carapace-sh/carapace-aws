package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Finds SageMaker resources that match a search query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_searchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_searchCmd).Standalone()

		sagemaker_searchCmd.Flags().String("cross-account-filter-option", "", "A cross account filter option.")
		sagemaker_searchCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		sagemaker_searchCmd.Flags().String("next-token", "", "If more than `MaxResults` resources match the specified `SearchExpression`, the response includes a `NextToken`.")
		sagemaker_searchCmd.Flags().String("resource", "", "The name of the SageMaker resource to search for.")
		sagemaker_searchCmd.Flags().String("search-expression", "", "A Boolean conditional statement.")
		sagemaker_searchCmd.Flags().String("sort-by", "", "The name of the resource property used to sort the `SearchResults`.")
		sagemaker_searchCmd.Flags().String("sort-order", "", "How `SearchResults` are ordered.")
		sagemaker_searchCmd.Flags().String("visibility-conditions", "", "Limits the results of your search request to the resources that you can access.")
		sagemaker_searchCmd.MarkFlagRequired("resource")
	})
	sagemakerCmd.AddCommand(sagemaker_searchCmd)
}

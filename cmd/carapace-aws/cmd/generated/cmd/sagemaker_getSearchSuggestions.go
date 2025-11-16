package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_getSearchSuggestionsCmd = &cobra.Command{
	Use:   "get-search-suggestions",
	Short: "An auto-complete API for the search functionality in the SageMaker console.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_getSearchSuggestionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_getSearchSuggestionsCmd).Standalone()

		sagemaker_getSearchSuggestionsCmd.Flags().String("resource", "", "The name of the SageMaker resource to search for.")
		sagemaker_getSearchSuggestionsCmd.Flags().String("suggestion-query", "", "Limits the property names that are included in the response.")
		sagemaker_getSearchSuggestionsCmd.MarkFlagRequired("resource")
	})
	sagemakerCmd.AddCommand(sagemaker_getSearchSuggestionsCmd)
}

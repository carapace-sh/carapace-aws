package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listModelMetadataCmd = &cobra.Command{
	Use:   "list-model-metadata",
	Short: "Lists the domain, framework, task, and model name of standard machine learning models found in common model zoos.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listModelMetadataCmd).Standalone()

	sagemaker_listModelMetadataCmd.Flags().String("max-results", "", "The maximum number of models to return in the response.")
	sagemaker_listModelMetadataCmd.Flags().String("next-token", "", "If the response to a previous `ListModelMetadataResponse` request was truncated, the response includes a NextToken.")
	sagemaker_listModelMetadataCmd.Flags().String("search-expression", "", "One or more filters that searches for the specified resource or resources in a search.")
	sagemakerCmd.AddCommand(sagemaker_listModelMetadataCmd)
}

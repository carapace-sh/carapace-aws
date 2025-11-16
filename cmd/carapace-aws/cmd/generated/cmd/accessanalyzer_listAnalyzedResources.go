package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_listAnalyzedResourcesCmd = &cobra.Command{
	Use:   "list-analyzed-resources",
	Short: "Retrieves a list of resources of the specified type that have been analyzed by the specified analyzer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_listAnalyzedResourcesCmd).Standalone()

	accessanalyzer_listAnalyzedResourcesCmd.Flags().String("analyzer-arn", "", "The [ARN of the analyzer](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources) to retrieve a list of analyzed resources from.")
	accessanalyzer_listAnalyzedResourcesCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	accessanalyzer_listAnalyzedResourcesCmd.Flags().String("next-token", "", "A token used for pagination of results returned.")
	accessanalyzer_listAnalyzedResourcesCmd.Flags().String("resource-type", "", "The type of resource.")
	accessanalyzer_listAnalyzedResourcesCmd.MarkFlagRequired("analyzer-arn")
	accessanalyzerCmd.AddCommand(accessanalyzer_listAnalyzedResourcesCmd)
}

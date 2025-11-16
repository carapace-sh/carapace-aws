package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_getAnalyzedResourceCmd = &cobra.Command{
	Use:   "get-analyzed-resource",
	Short: "Retrieves information about a resource that was analyzed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_getAnalyzedResourceCmd).Standalone()

	accessanalyzer_getAnalyzedResourceCmd.Flags().String("analyzer-arn", "", "The [ARN of the analyzer](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources) to retrieve information from.")
	accessanalyzer_getAnalyzedResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to retrieve information about.")
	accessanalyzer_getAnalyzedResourceCmd.MarkFlagRequired("analyzer-arn")
	accessanalyzer_getAnalyzedResourceCmd.MarkFlagRequired("resource-arn")
	accessanalyzerCmd.AddCommand(accessanalyzer_getAnalyzedResourceCmd)
}

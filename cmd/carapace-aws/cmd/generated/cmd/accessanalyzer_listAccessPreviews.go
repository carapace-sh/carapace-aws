package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_listAccessPreviewsCmd = &cobra.Command{
	Use:   "list-access-previews",
	Short: "Retrieves a list of access previews for the specified analyzer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_listAccessPreviewsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(accessanalyzer_listAccessPreviewsCmd).Standalone()

		accessanalyzer_listAccessPreviewsCmd.Flags().String("analyzer-arn", "", "The [ARN of the analyzer](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources) used to generate the access preview.")
		accessanalyzer_listAccessPreviewsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		accessanalyzer_listAccessPreviewsCmd.Flags().String("next-token", "", "A token used for pagination of results returned.")
		accessanalyzer_listAccessPreviewsCmd.MarkFlagRequired("analyzer-arn")
	})
	accessanalyzerCmd.AddCommand(accessanalyzer_listAccessPreviewsCmd)
}

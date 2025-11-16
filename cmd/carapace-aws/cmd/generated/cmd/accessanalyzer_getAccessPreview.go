package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_getAccessPreviewCmd = &cobra.Command{
	Use:   "get-access-preview",
	Short: "Retrieves information about an access preview for the specified analyzer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_getAccessPreviewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(accessanalyzer_getAccessPreviewCmd).Standalone()

		accessanalyzer_getAccessPreviewCmd.Flags().String("access-preview-id", "", "The unique ID for the access preview.")
		accessanalyzer_getAccessPreviewCmd.Flags().String("analyzer-arn", "", "The [ARN of the analyzer](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources) used to generate the access preview.")
		accessanalyzer_getAccessPreviewCmd.MarkFlagRequired("access-preview-id")
		accessanalyzer_getAccessPreviewCmd.MarkFlagRequired("analyzer-arn")
	})
	accessanalyzerCmd.AddCommand(accessanalyzer_getAccessPreviewCmd)
}

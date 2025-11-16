package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_getFindingV2Cmd = &cobra.Command{
	Use:   "get-finding-v2",
	Short: "Retrieves information about the specified finding.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_getFindingV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(accessanalyzer_getFindingV2Cmd).Standalone()

		accessanalyzer_getFindingV2Cmd.Flags().String("analyzer-arn", "", "The [ARN of the analyzer](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources) that generated the finding.")
		accessanalyzer_getFindingV2Cmd.Flags().String("id", "", "The ID of the finding to retrieve.")
		accessanalyzer_getFindingV2Cmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		accessanalyzer_getFindingV2Cmd.Flags().String("next-token", "", "A token used for pagination of results returned.")
		accessanalyzer_getFindingV2Cmd.MarkFlagRequired("analyzer-arn")
		accessanalyzer_getFindingV2Cmd.MarkFlagRequired("id")
	})
	accessanalyzerCmd.AddCommand(accessanalyzer_getFindingV2Cmd)
}

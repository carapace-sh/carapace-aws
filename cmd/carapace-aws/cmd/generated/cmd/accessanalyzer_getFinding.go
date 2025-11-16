package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_getFindingCmd = &cobra.Command{
	Use:   "get-finding",
	Short: "Retrieves information about the specified finding.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_getFindingCmd).Standalone()

	accessanalyzer_getFindingCmd.Flags().String("analyzer-arn", "", "The [ARN of the analyzer](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources) that generated the finding.")
	accessanalyzer_getFindingCmd.Flags().String("id", "", "The ID of the finding to retrieve.")
	accessanalyzer_getFindingCmd.MarkFlagRequired("analyzer-arn")
	accessanalyzer_getFindingCmd.MarkFlagRequired("id")
	accessanalyzerCmd.AddCommand(accessanalyzer_getFindingCmd)
}

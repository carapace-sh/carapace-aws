package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_getFindingsStatisticsCmd = &cobra.Command{
	Use:   "get-findings-statistics",
	Short: "Retrieves a list of aggregated finding statistics for an external access or unused access analyzer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_getFindingsStatisticsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(accessanalyzer_getFindingsStatisticsCmd).Standalone()

		accessanalyzer_getFindingsStatisticsCmd.Flags().String("analyzer-arn", "", "The [ARN of the analyzer](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources) used to generate the statistics.")
		accessanalyzer_getFindingsStatisticsCmd.MarkFlagRequired("analyzer-arn")
	})
	accessanalyzerCmd.AddCommand(accessanalyzer_getFindingsStatisticsCmd)
}

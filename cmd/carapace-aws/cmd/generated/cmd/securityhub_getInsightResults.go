package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_getInsightResultsCmd = &cobra.Command{
	Use:   "get-insight-results",
	Short: "Lists the results of the Security Hub insight specified by the insight ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_getInsightResultsCmd).Standalone()

	securityhub_getInsightResultsCmd.Flags().String("insight-arn", "", "The ARN of the insight for which to return results.")
	securityhub_getInsightResultsCmd.MarkFlagRequired("insight-arn")
	securityhubCmd.AddCommand(securityhub_getInsightResultsCmd)
}

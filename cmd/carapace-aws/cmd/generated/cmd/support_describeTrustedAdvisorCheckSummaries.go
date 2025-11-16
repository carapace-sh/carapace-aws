package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var support_describeTrustedAdvisorCheckSummariesCmd = &cobra.Command{
	Use:   "describe-trusted-advisor-check-summaries",
	Short: "Returns the results for the Trusted Advisor check summaries for the check IDs that you specified.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(support_describeTrustedAdvisorCheckSummariesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(support_describeTrustedAdvisorCheckSummariesCmd).Standalone()

		support_describeTrustedAdvisorCheckSummariesCmd.Flags().String("check-ids", "", "The IDs of the Trusted Advisor checks.")
		support_describeTrustedAdvisorCheckSummariesCmd.MarkFlagRequired("check-ids")
	})
	supportCmd.AddCommand(support_describeTrustedAdvisorCheckSummariesCmd)
}

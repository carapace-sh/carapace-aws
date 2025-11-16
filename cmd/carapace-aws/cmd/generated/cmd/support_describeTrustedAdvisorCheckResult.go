package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var support_describeTrustedAdvisorCheckResultCmd = &cobra.Command{
	Use:   "describe-trusted-advisor-check-result",
	Short: "Returns the results of the Trusted Advisor check that has the specified check ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(support_describeTrustedAdvisorCheckResultCmd).Standalone()

	support_describeTrustedAdvisorCheckResultCmd.Flags().String("check-id", "", "The unique identifier for the Trusted Advisor check.")
	support_describeTrustedAdvisorCheckResultCmd.Flags().String("language", "", "The ISO 639-1 code for the language that you want your check results to appear in.")
	support_describeTrustedAdvisorCheckResultCmd.MarkFlagRequired("check-id")
	supportCmd.AddCommand(support_describeTrustedAdvisorCheckResultCmd)
}

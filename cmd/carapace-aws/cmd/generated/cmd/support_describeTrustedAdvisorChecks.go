package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var support_describeTrustedAdvisorChecksCmd = &cobra.Command{
	Use:   "describe-trusted-advisor-checks",
	Short: "Returns information about all available Trusted Advisor checks, including the name, ID, category, description, and metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(support_describeTrustedAdvisorChecksCmd).Standalone()

	support_describeTrustedAdvisorChecksCmd.Flags().String("language", "", "The ISO 639-1 code for the language that you want your checks to appear in.")
	support_describeTrustedAdvisorChecksCmd.MarkFlagRequired("language")
	supportCmd.AddCommand(support_describeTrustedAdvisorChecksCmd)
}

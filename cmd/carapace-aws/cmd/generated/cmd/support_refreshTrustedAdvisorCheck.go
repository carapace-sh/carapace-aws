package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var support_refreshTrustedAdvisorCheckCmd = &cobra.Command{
	Use:   "refresh-trusted-advisor-check",
	Short: "Refreshes the Trusted Advisor check that you specify using the check ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(support_refreshTrustedAdvisorCheckCmd).Standalone()

	support_refreshTrustedAdvisorCheckCmd.Flags().String("check-id", "", "The unique identifier for the Trusted Advisor check to refresh.")
	support_refreshTrustedAdvisorCheckCmd.MarkFlagRequired("check-id")
	supportCmd.AddCommand(support_refreshTrustedAdvisorCheckCmd)
}

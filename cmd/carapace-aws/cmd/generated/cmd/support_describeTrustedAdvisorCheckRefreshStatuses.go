package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var support_describeTrustedAdvisorCheckRefreshStatusesCmd = &cobra.Command{
	Use:   "describe-trusted-advisor-check-refresh-statuses",
	Short: "Returns the refresh status of the Trusted Advisor checks that have the specified check IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(support_describeTrustedAdvisorCheckRefreshStatusesCmd).Standalone()

	support_describeTrustedAdvisorCheckRefreshStatusesCmd.Flags().String("check-ids", "", "The IDs of the Trusted Advisor checks to get the status.")
	support_describeTrustedAdvisorCheckRefreshStatusesCmd.MarkFlagRequired("check-ids")
	supportCmd.AddCommand(support_describeTrustedAdvisorCheckRefreshStatusesCmd)
}

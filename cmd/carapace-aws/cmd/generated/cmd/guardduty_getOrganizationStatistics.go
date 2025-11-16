package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_getOrganizationStatisticsCmd = &cobra.Command{
	Use:   "get-organization-statistics",
	Short: "Retrieves how many active member accounts have each feature enabled within GuardDuty.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_getOrganizationStatisticsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_getOrganizationStatisticsCmd).Standalone()

	})
	guarddutyCmd.AddCommand(guardduty_getOrganizationStatisticsCmd)
}

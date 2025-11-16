package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_getMembersCmd = &cobra.Command{
	Use:   "get-members",
	Short: "Retrieves GuardDuty member accounts (of the current GuardDuty administrator account) specified by the account IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_getMembersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_getMembersCmd).Standalone()

		guardduty_getMembersCmd.Flags().String("account-ids", "", "A list of account IDs of the GuardDuty member accounts that you want to describe.")
		guardduty_getMembersCmd.Flags().String("detector-id", "", "The unique ID of the detector of the GuardDuty account whose members you want to retrieve.")
		guardduty_getMembersCmd.MarkFlagRequired("account-ids")
		guardduty_getMembersCmd.MarkFlagRequired("detector-id")
	})
	guarddutyCmd.AddCommand(guardduty_getMembersCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getProfileHistoryRecordCmd = &cobra.Command{
	Use:   "get-profile-history-record",
	Short: "Returns a history record for a specific profile, for a specific domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getProfileHistoryRecordCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_getProfileHistoryRecordCmd).Standalone()

		customerProfiles_getProfileHistoryRecordCmd.Flags().String("domain-name", "", "The unique name of the domain for which to return a profile history record.")
		customerProfiles_getProfileHistoryRecordCmd.Flags().String("id", "", "The unique identifier of the profile history record to return.")
		customerProfiles_getProfileHistoryRecordCmd.Flags().String("profile-id", "", "The unique identifier of the profile for which to return a history record.")
		customerProfiles_getProfileHistoryRecordCmd.MarkFlagRequired("domain-name")
		customerProfiles_getProfileHistoryRecordCmd.MarkFlagRequired("id")
		customerProfiles_getProfileHistoryRecordCmd.MarkFlagRequired("profile-id")
	})
	customerProfilesCmd.AddCommand(customerProfiles_getProfileHistoryRecordCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_listProfileHistoryRecordsCmd = &cobra.Command{
	Use:   "list-profile-history-records",
	Short: "Returns a list of history records for a specific profile, for a specific domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_listProfileHistoryRecordsCmd).Standalone()

	customerProfiles_listProfileHistoryRecordsCmd.Flags().String("action-type", "", "Applies a filter to include profile history records only with the specified `ActionType` value in the response.")
	customerProfiles_listProfileHistoryRecordsCmd.Flags().String("domain-name", "", "The unique name of the domain for which to return profile history records.")
	customerProfiles_listProfileHistoryRecordsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	customerProfiles_listProfileHistoryRecordsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	customerProfiles_listProfileHistoryRecordsCmd.Flags().String("object-type-name", "", "Applies a filter to include profile history records only with the specified `ObjectTypeName` value in the response.")
	customerProfiles_listProfileHistoryRecordsCmd.Flags().String("performed-by", "", "Applies a filter to include profile history records only with the specified `PerformedBy` value in the response.")
	customerProfiles_listProfileHistoryRecordsCmd.Flags().String("profile-id", "", "The identifier of the profile to be taken.")
	customerProfiles_listProfileHistoryRecordsCmd.MarkFlagRequired("domain-name")
	customerProfiles_listProfileHistoryRecordsCmd.MarkFlagRequired("profile-id")
	customerProfilesCmd.AddCommand(customerProfiles_listProfileHistoryRecordsCmd)
}

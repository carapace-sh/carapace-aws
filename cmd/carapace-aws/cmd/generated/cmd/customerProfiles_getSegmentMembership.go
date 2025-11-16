package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getSegmentMembershipCmd = &cobra.Command{
	Use:   "get-segment-membership",
	Short: "Determines if the given profiles are within a segment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getSegmentMembershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_getSegmentMembershipCmd).Standalone()

		customerProfiles_getSegmentMembershipCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_getSegmentMembershipCmd.Flags().String("profile-ids", "", "The list of profile IDs to query for.")
		customerProfiles_getSegmentMembershipCmd.Flags().String("segment-definition-name", "", "The Id of the wanted segment.")
		customerProfiles_getSegmentMembershipCmd.MarkFlagRequired("domain-name")
		customerProfiles_getSegmentMembershipCmd.MarkFlagRequired("profile-ids")
		customerProfiles_getSegmentMembershipCmd.MarkFlagRequired("segment-definition-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_getSegmentMembershipCmd)
}

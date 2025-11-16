package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_mergeProfilesCmd = &cobra.Command{
	Use:   "merge-profiles",
	Short: "Runs an AWS Lambda job that does the following:\n\n1. All the profileKeys in the `ProfileToBeMerged` will be moved to the main profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_mergeProfilesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_mergeProfilesCmd).Standalone()

		customerProfiles_mergeProfilesCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_mergeProfilesCmd.Flags().String("field-source-profile-ids", "", "The identifiers of the fields in the profile that has the information you want to apply to the merge.")
		customerProfiles_mergeProfilesCmd.Flags().String("main-profile-id", "", "The identifier of the profile to be taken.")
		customerProfiles_mergeProfilesCmd.Flags().String("profile-ids-to-be-merged", "", "The identifier of the profile to be merged into MainProfileId.")
		customerProfiles_mergeProfilesCmd.MarkFlagRequired("domain-name")
		customerProfiles_mergeProfilesCmd.MarkFlagRequired("main-profile-id")
		customerProfiles_mergeProfilesCmd.MarkFlagRequired("profile-ids-to-be-merged")
	})
	customerProfilesCmd.AddCommand(customerProfiles_mergeProfilesCmd)
}

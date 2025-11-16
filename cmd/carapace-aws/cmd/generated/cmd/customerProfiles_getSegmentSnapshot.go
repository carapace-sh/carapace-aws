package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getSegmentSnapshotCmd = &cobra.Command{
	Use:   "get-segment-snapshot",
	Short: "Retrieve the latest status of a segment snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getSegmentSnapshotCmd).Standalone()

	customerProfiles_getSegmentSnapshotCmd.Flags().String("domain-name", "", "The unique identifier of the domain.")
	customerProfiles_getSegmentSnapshotCmd.Flags().String("segment-definition-name", "", "The unique name of the segment definition.")
	customerProfiles_getSegmentSnapshotCmd.Flags().String("snapshot-id", "", "The unique identifier of the segment snapshot.")
	customerProfiles_getSegmentSnapshotCmd.MarkFlagRequired("domain-name")
	customerProfiles_getSegmentSnapshotCmd.MarkFlagRequired("segment-definition-name")
	customerProfiles_getSegmentSnapshotCmd.MarkFlagRequired("snapshot-id")
	customerProfilesCmd.AddCommand(customerProfiles_getSegmentSnapshotCmd)
}

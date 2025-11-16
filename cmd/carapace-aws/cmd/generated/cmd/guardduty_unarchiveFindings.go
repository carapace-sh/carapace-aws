package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_unarchiveFindingsCmd = &cobra.Command{
	Use:   "unarchive-findings",
	Short: "Unarchives GuardDuty findings specified by the `findingIds`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_unarchiveFindingsCmd).Standalone()

	guardduty_unarchiveFindingsCmd.Flags().String("detector-id", "", "The ID of the detector associated with the findings to unarchive.")
	guardduty_unarchiveFindingsCmd.Flags().String("finding-ids", "", "The IDs of the findings to unarchive.")
	guardduty_unarchiveFindingsCmd.MarkFlagRequired("detector-id")
	guardduty_unarchiveFindingsCmd.MarkFlagRequired("finding-ids")
	guarddutyCmd.AddCommand(guardduty_unarchiveFindingsCmd)
}

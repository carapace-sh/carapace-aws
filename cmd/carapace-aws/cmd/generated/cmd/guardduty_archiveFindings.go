package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_archiveFindingsCmd = &cobra.Command{
	Use:   "archive-findings",
	Short: "Archives GuardDuty findings that are specified by the list of finding IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_archiveFindingsCmd).Standalone()

	guardduty_archiveFindingsCmd.Flags().String("detector-id", "", "The ID of the detector that specifies the GuardDuty service whose findings you want to archive.")
	guardduty_archiveFindingsCmd.Flags().String("finding-ids", "", "The IDs of the findings that you want to archive.")
	guardduty_archiveFindingsCmd.MarkFlagRequired("detector-id")
	guardduty_archiveFindingsCmd.MarkFlagRequired("finding-ids")
	guarddutyCmd.AddCommand(guardduty_archiveFindingsCmd)
}

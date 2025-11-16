package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_getFindingsCmd = &cobra.Command{
	Use:   "get-findings",
	Short: "Describes Amazon GuardDuty findings specified by finding IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_getFindingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_getFindingsCmd).Standalone()

		guardduty_getFindingsCmd.Flags().String("detector-id", "", "The ID of the detector that specifies the GuardDuty service whose findings you want to retrieve.")
		guardduty_getFindingsCmd.Flags().String("finding-ids", "", "The IDs of the findings that you want to retrieve.")
		guardduty_getFindingsCmd.Flags().String("sort-criteria", "", "Represents the criteria used for sorting findings.")
		guardduty_getFindingsCmd.MarkFlagRequired("detector-id")
		guardduty_getFindingsCmd.MarkFlagRequired("finding-ids")
	})
	guarddutyCmd.AddCommand(guardduty_getFindingsCmd)
}

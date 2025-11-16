package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_updateFindingsFeedbackCmd = &cobra.Command{
	Use:   "update-findings-feedback",
	Short: "Marks the specified GuardDuty findings as useful or not useful.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_updateFindingsFeedbackCmd).Standalone()

	guardduty_updateFindingsFeedbackCmd.Flags().String("comments", "", "Additional feedback about the GuardDuty findings.")
	guardduty_updateFindingsFeedbackCmd.Flags().String("detector-id", "", "The ID of the detector that is associated with the findings for which you want to update the feedback.")
	guardduty_updateFindingsFeedbackCmd.Flags().String("feedback", "", "The feedback for the finding.")
	guardduty_updateFindingsFeedbackCmd.Flags().String("finding-ids", "", "The IDs of the findings that you want to mark as useful or not useful.")
	guardduty_updateFindingsFeedbackCmd.MarkFlagRequired("detector-id")
	guardduty_updateFindingsFeedbackCmd.MarkFlagRequired("feedback")
	guardduty_updateFindingsFeedbackCmd.MarkFlagRequired("finding-ids")
	guarddutyCmd.AddCommand(guardduty_updateFindingsFeedbackCmd)
}

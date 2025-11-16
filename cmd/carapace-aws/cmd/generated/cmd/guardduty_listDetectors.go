package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_listDetectorsCmd = &cobra.Command{
	Use:   "list-detectors",
	Short: "Lists detectorIds of all the existing Amazon GuardDuty detector resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_listDetectorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_listDetectorsCmd).Standalone()

		guardduty_listDetectorsCmd.Flags().String("max-results", "", "You can use this parameter to indicate the maximum number of items that you want in the response.")
		guardduty_listDetectorsCmd.Flags().String("next-token", "", "You can use this parameter when paginating results.")
	})
	guarddutyCmd.AddCommand(guardduty_listDetectorsCmd)
}

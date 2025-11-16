package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_getMemberDetectorsCmd = &cobra.Command{
	Use:   "get-member-detectors",
	Short: "Describes which data sources are enabled for the member account's detector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_getMemberDetectorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_getMemberDetectorsCmd).Standalone()

		guardduty_getMemberDetectorsCmd.Flags().String("account-ids", "", "A list of member account IDs.")
		guardduty_getMemberDetectorsCmd.Flags().String("detector-id", "", "The detector ID for the administrator account.")
		guardduty_getMemberDetectorsCmd.MarkFlagRequired("account-ids")
		guardduty_getMemberDetectorsCmd.MarkFlagRequired("detector-id")
	})
	guarddutyCmd.AddCommand(guardduty_getMemberDetectorsCmd)
}

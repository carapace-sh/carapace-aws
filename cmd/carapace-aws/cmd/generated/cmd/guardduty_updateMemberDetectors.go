package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_updateMemberDetectorsCmd = &cobra.Command{
	Use:   "update-member-detectors",
	Short: "Contains information on member accounts to be updated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_updateMemberDetectorsCmd).Standalone()

	guardduty_updateMemberDetectorsCmd.Flags().String("account-ids", "", "A list of member account IDs to be updated.")
	guardduty_updateMemberDetectorsCmd.Flags().String("data-sources", "", "Describes which data sources will be updated.")
	guardduty_updateMemberDetectorsCmd.Flags().String("detector-id", "", "The detector ID of the administrator account.")
	guardduty_updateMemberDetectorsCmd.Flags().String("features", "", "A list of features that will be updated for the specified member accounts.")
	guardduty_updateMemberDetectorsCmd.MarkFlagRequired("account-ids")
	guardduty_updateMemberDetectorsCmd.MarkFlagRequired("detector-id")
	guarddutyCmd.AddCommand(guardduty_updateMemberDetectorsCmd)
}

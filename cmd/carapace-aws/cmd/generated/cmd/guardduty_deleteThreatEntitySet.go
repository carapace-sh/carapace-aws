package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_deleteThreatEntitySetCmd = &cobra.Command{
	Use:   "delete-threat-entity-set",
	Short: "Deletes the threat entity set that is associated with the specified `threatEntitySetId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_deleteThreatEntitySetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_deleteThreatEntitySetCmd).Standalone()

		guardduty_deleteThreatEntitySetCmd.Flags().String("detector-id", "", "The unique ID of the detector associated with the threat entity set resource.")
		guardduty_deleteThreatEntitySetCmd.Flags().String("threat-entity-set-id", "", "The unique ID that helps GuardDuty identify which threat entity set needs to be deleted.")
		guardduty_deleteThreatEntitySetCmd.MarkFlagRequired("detector-id")
		guardduty_deleteThreatEntitySetCmd.MarkFlagRequired("threat-entity-set-id")
	})
	guarddutyCmd.AddCommand(guardduty_deleteThreatEntitySetCmd)
}

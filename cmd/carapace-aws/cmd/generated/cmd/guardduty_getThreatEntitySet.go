package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_getThreatEntitySetCmd = &cobra.Command{
	Use:   "get-threat-entity-set",
	Short: "Retrieves the threat entity set associated with the specified `threatEntitySetId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_getThreatEntitySetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_getThreatEntitySetCmd).Standalone()

		guardduty_getThreatEntitySetCmd.Flags().String("detector-id", "", "The unique ID of the detector associated with the threat entity set resource.")
		guardduty_getThreatEntitySetCmd.Flags().String("threat-entity-set-id", "", "The unique ID that helps GuardDuty identify the threat entity set.")
		guardduty_getThreatEntitySetCmd.MarkFlagRequired("detector-id")
		guardduty_getThreatEntitySetCmd.MarkFlagRequired("threat-entity-set-id")
	})
	guarddutyCmd.AddCommand(guardduty_getThreatEntitySetCmd)
}

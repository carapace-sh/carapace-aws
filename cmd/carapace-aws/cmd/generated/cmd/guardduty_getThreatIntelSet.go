package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_getThreatIntelSetCmd = &cobra.Command{
	Use:   "get-threat-intel-set",
	Short: "Retrieves the ThreatIntelSet that is specified by the ThreatIntelSet ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_getThreatIntelSetCmd).Standalone()

	guardduty_getThreatIntelSetCmd.Flags().String("detector-id", "", "The unique ID of the detector that is associated with the threatIntelSet.")
	guardduty_getThreatIntelSetCmd.Flags().String("threat-intel-set-id", "", "The unique ID of the threatIntelSet that you want to get.")
	guardduty_getThreatIntelSetCmd.MarkFlagRequired("detector-id")
	guardduty_getThreatIntelSetCmd.MarkFlagRequired("threat-intel-set-id")
	guarddutyCmd.AddCommand(guardduty_getThreatIntelSetCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_deleteThreatIntelSetCmd = &cobra.Command{
	Use:   "delete-threat-intel-set",
	Short: "Deletes the ThreatIntelSet specified by the ThreatIntelSet ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_deleteThreatIntelSetCmd).Standalone()

	guardduty_deleteThreatIntelSetCmd.Flags().String("detector-id", "", "The unique ID of the detector that is associated with the threatIntelSet.")
	guardduty_deleteThreatIntelSetCmd.Flags().String("threat-intel-set-id", "", "The unique ID of the threatIntelSet that you want to delete.")
	guardduty_deleteThreatIntelSetCmd.MarkFlagRequired("detector-id")
	guardduty_deleteThreatIntelSetCmd.MarkFlagRequired("threat-intel-set-id")
	guarddutyCmd.AddCommand(guardduty_deleteThreatIntelSetCmd)
}

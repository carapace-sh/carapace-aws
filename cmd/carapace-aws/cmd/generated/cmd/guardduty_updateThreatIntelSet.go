package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_updateThreatIntelSetCmd = &cobra.Command{
	Use:   "update-threat-intel-set",
	Short: "Updates the ThreatIntelSet specified by the ThreatIntelSet ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_updateThreatIntelSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_updateThreatIntelSetCmd).Standalone()

		guardduty_updateThreatIntelSetCmd.Flags().Bool("activate", false, "The updated Boolean value that specifies whether the ThreateIntelSet is active or not.")
		guardduty_updateThreatIntelSetCmd.Flags().String("detector-id", "", "The detectorID that specifies the GuardDuty service whose ThreatIntelSet you want to update.")
		guardduty_updateThreatIntelSetCmd.Flags().String("expected-bucket-owner", "", "The Amazon Web Services account ID that owns the Amazon S3 bucket specified in the **location** parameter.")
		guardduty_updateThreatIntelSetCmd.Flags().String("location", "", "The updated URI of the file that contains the ThreateIntelSet.")
		guardduty_updateThreatIntelSetCmd.Flags().String("name", "", "The unique ID that specifies the ThreatIntelSet that you want to update.")
		guardduty_updateThreatIntelSetCmd.Flags().Bool("no-activate", false, "The updated Boolean value that specifies whether the ThreateIntelSet is active or not.")
		guardduty_updateThreatIntelSetCmd.Flags().String("threat-intel-set-id", "", "The unique ID that specifies the ThreatIntelSet that you want to update.")
		guardduty_updateThreatIntelSetCmd.MarkFlagRequired("detector-id")
		guardduty_updateThreatIntelSetCmd.Flag("no-activate").Hidden = true
		guardduty_updateThreatIntelSetCmd.MarkFlagRequired("threat-intel-set-id")
	})
	guarddutyCmd.AddCommand(guardduty_updateThreatIntelSetCmd)
}

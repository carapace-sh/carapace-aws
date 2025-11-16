package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_updateDetectorCmd = &cobra.Command{
	Use:   "update-detector",
	Short: "Updates the GuardDuty detector specified by the detector ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_updateDetectorCmd).Standalone()

	guardduty_updateDetectorCmd.Flags().String("data-sources", "", "Describes which data sources will be updated.")
	guardduty_updateDetectorCmd.Flags().String("detector-id", "", "The unique ID of the detector to update.")
	guardduty_updateDetectorCmd.Flags().Bool("enable", false, "Specifies whether the detector is enabled or not enabled.")
	guardduty_updateDetectorCmd.Flags().String("features", "", "Provides the features that will be updated for the detector.")
	guardduty_updateDetectorCmd.Flags().String("finding-publishing-frequency", "", "An enum value that specifies how frequently findings are exported, such as to CloudWatch Events.")
	guardduty_updateDetectorCmd.Flags().Bool("no-enable", false, "Specifies whether the detector is enabled or not enabled.")
	guardduty_updateDetectorCmd.MarkFlagRequired("detector-id")
	guardduty_updateDetectorCmd.Flag("no-enable").Hidden = true
	guarddutyCmd.AddCommand(guardduty_updateDetectorCmd)
}

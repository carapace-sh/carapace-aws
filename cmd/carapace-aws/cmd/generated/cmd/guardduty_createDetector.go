package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_createDetectorCmd = &cobra.Command{
	Use:   "create-detector",
	Short: "Creates a single GuardDuty detector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_createDetectorCmd).Standalone()

	guardduty_createDetectorCmd.Flags().String("client-token", "", "The idempotency token for the create request.")
	guardduty_createDetectorCmd.Flags().String("data-sources", "", "Describes which data sources will be enabled for the detector.")
	guardduty_createDetectorCmd.Flags().Bool("enable", false, "A Boolean value that specifies whether the detector is to be enabled.")
	guardduty_createDetectorCmd.Flags().String("features", "", "A list of features that will be configured for the detector.")
	guardduty_createDetectorCmd.Flags().String("finding-publishing-frequency", "", "A value that specifies how frequently updated findings are exported.")
	guardduty_createDetectorCmd.Flags().Bool("no-enable", false, "A Boolean value that specifies whether the detector is to be enabled.")
	guardduty_createDetectorCmd.Flags().String("tags", "", "The tags to be added to a new detector resource.")
	guardduty_createDetectorCmd.MarkFlagRequired("enable")
	guardduty_createDetectorCmd.Flag("no-enable").Hidden = true
	guardduty_createDetectorCmd.MarkFlagRequired("no-enable")
	guarddutyCmd.AddCommand(guardduty_createDetectorCmd)
}

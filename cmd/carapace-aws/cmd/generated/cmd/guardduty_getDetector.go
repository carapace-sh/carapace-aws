package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_getDetectorCmd = &cobra.Command{
	Use:   "get-detector",
	Short: "Retrieves a GuardDuty detector specified by the detectorId.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_getDetectorCmd).Standalone()

	guardduty_getDetectorCmd.Flags().String("detector-id", "", "The unique ID of the detector that you want to get.")
	guardduty_getDetectorCmd.MarkFlagRequired("detector-id")
	guarddutyCmd.AddCommand(guardduty_getDetectorCmd)
}

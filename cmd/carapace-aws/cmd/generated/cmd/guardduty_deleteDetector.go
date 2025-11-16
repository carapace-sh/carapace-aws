package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_deleteDetectorCmd = &cobra.Command{
	Use:   "delete-detector",
	Short: "Deletes an Amazon GuardDuty detector that is specified by the detector ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_deleteDetectorCmd).Standalone()

	guardduty_deleteDetectorCmd.Flags().String("detector-id", "", "The unique ID of the detector that you want to delete.")
	guardduty_deleteDetectorCmd.MarkFlagRequired("detector-id")
	guarddutyCmd.AddCommand(guardduty_deleteDetectorCmd)
}

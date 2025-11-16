package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_updateDetectorVersionStatusCmd = &cobra.Command{
	Use:   "update-detector-version-status",
	Short: "Updates the detector version’s status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_updateDetectorVersionStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_updateDetectorVersionStatusCmd).Standalone()

		frauddetector_updateDetectorVersionStatusCmd.Flags().String("detector-id", "", "The detector ID.")
		frauddetector_updateDetectorVersionStatusCmd.Flags().String("detector-version-id", "", "The detector version ID.")
		frauddetector_updateDetectorVersionStatusCmd.Flags().String("status", "", "The new status.")
		frauddetector_updateDetectorVersionStatusCmd.MarkFlagRequired("detector-id")
		frauddetector_updateDetectorVersionStatusCmd.MarkFlagRequired("detector-version-id")
		frauddetector_updateDetectorVersionStatusCmd.MarkFlagRequired("status")
	})
	frauddetectorCmd.AddCommand(frauddetector_updateDetectorVersionStatusCmd)
}

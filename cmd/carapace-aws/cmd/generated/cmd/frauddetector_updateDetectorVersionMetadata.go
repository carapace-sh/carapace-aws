package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_updateDetectorVersionMetadataCmd = &cobra.Command{
	Use:   "update-detector-version-metadata",
	Short: "Updates the detector version's description.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_updateDetectorVersionMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_updateDetectorVersionMetadataCmd).Standalone()

		frauddetector_updateDetectorVersionMetadataCmd.Flags().String("description", "", "The description.")
		frauddetector_updateDetectorVersionMetadataCmd.Flags().String("detector-id", "", "The detector ID.")
		frauddetector_updateDetectorVersionMetadataCmd.Flags().String("detector-version-id", "", "The detector version ID.")
		frauddetector_updateDetectorVersionMetadataCmd.MarkFlagRequired("description")
		frauddetector_updateDetectorVersionMetadataCmd.MarkFlagRequired("detector-id")
		frauddetector_updateDetectorVersionMetadataCmd.MarkFlagRequired("detector-version-id")
	})
	frauddetectorCmd.AddCommand(frauddetector_updateDetectorVersionMetadataCmd)
}

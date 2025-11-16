package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_updateDetectorVersionCmd = &cobra.Command{
	Use:   "update-detector-version",
	Short: "Updates a detector version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_updateDetectorVersionCmd).Standalone()

	frauddetector_updateDetectorVersionCmd.Flags().String("description", "", "The detector version description.")
	frauddetector_updateDetectorVersionCmd.Flags().String("detector-id", "", "The parent detector ID for the detector version you want to update.")
	frauddetector_updateDetectorVersionCmd.Flags().String("detector-version-id", "", "The detector version ID.")
	frauddetector_updateDetectorVersionCmd.Flags().String("external-model-endpoints", "", "The Amazon SageMaker model endpoints to include in the detector version.")
	frauddetector_updateDetectorVersionCmd.Flags().String("model-versions", "", "The model versions to include in the detector version.")
	frauddetector_updateDetectorVersionCmd.Flags().String("rule-execution-mode", "", "The rule execution mode to add to the detector.")
	frauddetector_updateDetectorVersionCmd.Flags().String("rules", "", "The rules to include in the detector version.")
	frauddetector_updateDetectorVersionCmd.MarkFlagRequired("detector-id")
	frauddetector_updateDetectorVersionCmd.MarkFlagRequired("detector-version-id")
	frauddetector_updateDetectorVersionCmd.MarkFlagRequired("external-model-endpoints")
	frauddetector_updateDetectorVersionCmd.MarkFlagRequired("rules")
	frauddetectorCmd.AddCommand(frauddetector_updateDetectorVersionCmd)
}

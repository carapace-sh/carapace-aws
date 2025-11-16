package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_createDetectorVersionCmd = &cobra.Command{
	Use:   "create-detector-version",
	Short: "Creates a detector version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_createDetectorVersionCmd).Standalone()

	frauddetector_createDetectorVersionCmd.Flags().String("description", "", "The description of the detector version.")
	frauddetector_createDetectorVersionCmd.Flags().String("detector-id", "", "The ID of the detector under which you want to create a new version.")
	frauddetector_createDetectorVersionCmd.Flags().String("external-model-endpoints", "", "The Amazon Sagemaker model endpoints to include in the detector version.")
	frauddetector_createDetectorVersionCmd.Flags().String("model-versions", "", "The model versions to include in the detector version.")
	frauddetector_createDetectorVersionCmd.Flags().String("rule-execution-mode", "", "The rule execution mode for the rules included in the detector version.")
	frauddetector_createDetectorVersionCmd.Flags().String("rules", "", "The rules to include in the detector version.")
	frauddetector_createDetectorVersionCmd.Flags().String("tags", "", "A collection of key and value pairs.")
	frauddetector_createDetectorVersionCmd.MarkFlagRequired("detector-id")
	frauddetector_createDetectorVersionCmd.MarkFlagRequired("rules")
	frauddetectorCmd.AddCommand(frauddetector_createDetectorVersionCmd)
}

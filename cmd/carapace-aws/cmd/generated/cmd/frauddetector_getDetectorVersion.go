package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_getDetectorVersionCmd = &cobra.Command{
	Use:   "get-detector-version",
	Short: "Gets a particular detector version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_getDetectorVersionCmd).Standalone()

	frauddetector_getDetectorVersionCmd.Flags().String("detector-id", "", "The detector ID.")
	frauddetector_getDetectorVersionCmd.Flags().String("detector-version-id", "", "The detector version ID.")
	frauddetector_getDetectorVersionCmd.MarkFlagRequired("detector-id")
	frauddetector_getDetectorVersionCmd.MarkFlagRequired("detector-version-id")
	frauddetectorCmd.AddCommand(frauddetector_getDetectorVersionCmd)
}

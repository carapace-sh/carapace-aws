package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_deleteDetectorVersionCmd = &cobra.Command{
	Use:   "delete-detector-version",
	Short: "Deletes the detector version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_deleteDetectorVersionCmd).Standalone()

	frauddetector_deleteDetectorVersionCmd.Flags().String("detector-id", "", "The ID of the parent detector for the detector version to delete.")
	frauddetector_deleteDetectorVersionCmd.Flags().String("detector-version-id", "", "The ID of the detector version to delete.")
	frauddetector_deleteDetectorVersionCmd.MarkFlagRequired("detector-id")
	frauddetector_deleteDetectorVersionCmd.MarkFlagRequired("detector-version-id")
	frauddetectorCmd.AddCommand(frauddetector_deleteDetectorVersionCmd)
}

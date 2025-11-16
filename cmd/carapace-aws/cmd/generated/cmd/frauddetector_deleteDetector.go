package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_deleteDetectorCmd = &cobra.Command{
	Use:   "delete-detector",
	Short: "Deletes the detector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_deleteDetectorCmd).Standalone()

	frauddetector_deleteDetectorCmd.Flags().String("detector-id", "", "The ID of the detector to delete.")
	frauddetector_deleteDetectorCmd.MarkFlagRequired("detector-id")
	frauddetectorCmd.AddCommand(frauddetector_deleteDetectorCmd)
}

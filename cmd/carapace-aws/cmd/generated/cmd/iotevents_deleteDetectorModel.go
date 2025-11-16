package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_deleteDetectorModelCmd = &cobra.Command{
	Use:   "delete-detector-model",
	Short: "Deletes a detector model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_deleteDetectorModelCmd).Standalone()

	iotevents_deleteDetectorModelCmd.Flags().String("detector-model-name", "", "The name of the detector model to be deleted.")
	iotevents_deleteDetectorModelCmd.MarkFlagRequired("detector-model-name")
	ioteventsCmd.AddCommand(iotevents_deleteDetectorModelCmd)
}

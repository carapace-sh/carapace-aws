package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_describeDetectorModelCmd = &cobra.Command{
	Use:   "describe-detector-model",
	Short: "Describes a detector model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_describeDetectorModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotevents_describeDetectorModelCmd).Standalone()

		iotevents_describeDetectorModelCmd.Flags().String("detector-model-name", "", "The name of the detector model.")
		iotevents_describeDetectorModelCmd.Flags().String("detector-model-version", "", "The version of the detector model.")
		iotevents_describeDetectorModelCmd.MarkFlagRequired("detector-model-name")
	})
	ioteventsCmd.AddCommand(iotevents_describeDetectorModelCmd)
}

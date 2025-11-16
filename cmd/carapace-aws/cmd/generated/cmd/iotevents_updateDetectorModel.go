package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_updateDetectorModelCmd = &cobra.Command{
	Use:   "update-detector-model",
	Short: "Updates a detector model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_updateDetectorModelCmd).Standalone()

	iotevents_updateDetectorModelCmd.Flags().String("detector-model-definition", "", "Information that defines how a detector operates.")
	iotevents_updateDetectorModelCmd.Flags().String("detector-model-description", "", "A brief description of the detector model.")
	iotevents_updateDetectorModelCmd.Flags().String("detector-model-name", "", "The name of the detector model that is updated.")
	iotevents_updateDetectorModelCmd.Flags().String("evaluation-method", "", "Information about the order in which events are evaluated and how actions are executed.")
	iotevents_updateDetectorModelCmd.Flags().String("role-arn", "", "The ARN of the role that grants permission to AWS IoT Events to perform its operations.")
	iotevents_updateDetectorModelCmd.MarkFlagRequired("detector-model-definition")
	iotevents_updateDetectorModelCmd.MarkFlagRequired("detector-model-name")
	iotevents_updateDetectorModelCmd.MarkFlagRequired("role-arn")
	ioteventsCmd.AddCommand(iotevents_updateDetectorModelCmd)
}

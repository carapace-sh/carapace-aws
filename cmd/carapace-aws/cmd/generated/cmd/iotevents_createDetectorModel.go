package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_createDetectorModelCmd = &cobra.Command{
	Use:   "create-detector-model",
	Short: "Creates a detector model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_createDetectorModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotevents_createDetectorModelCmd).Standalone()

		iotevents_createDetectorModelCmd.Flags().String("detector-model-definition", "", "Information that defines how the detectors operate.")
		iotevents_createDetectorModelCmd.Flags().String("detector-model-description", "", "A brief description of the detector model.")
		iotevents_createDetectorModelCmd.Flags().String("detector-model-name", "", "The name of the detector model.")
		iotevents_createDetectorModelCmd.Flags().String("evaluation-method", "", "Information about the order in which events are evaluated and how actions are executed.")
		iotevents_createDetectorModelCmd.Flags().String("key", "", "The input attribute key used to identify a device or system to create a detector (an instance of the detector model) and then to route each input received to the appropriate detector (instance).")
		iotevents_createDetectorModelCmd.Flags().String("role-arn", "", "The ARN of the role that grants permission to AWS IoT Events to perform its operations.")
		iotevents_createDetectorModelCmd.Flags().String("tags", "", "Metadata that can be used to manage the detector model.")
		iotevents_createDetectorModelCmd.MarkFlagRequired("detector-model-definition")
		iotevents_createDetectorModelCmd.MarkFlagRequired("detector-model-name")
		iotevents_createDetectorModelCmd.MarkFlagRequired("role-arn")
	})
	ioteventsCmd.AddCommand(iotevents_createDetectorModelCmd)
}

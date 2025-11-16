package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_describeInferenceSchedulerCmd = &cobra.Command{
	Use:   "describe-inference-scheduler",
	Short: "Specifies information about the inference scheduler being used, including name, model, status, and associated metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_describeInferenceSchedulerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_describeInferenceSchedulerCmd).Standalone()

		lookoutequipment_describeInferenceSchedulerCmd.Flags().String("inference-scheduler-name", "", "The name of the inference scheduler being described.")
		lookoutequipment_describeInferenceSchedulerCmd.MarkFlagRequired("inference-scheduler-name")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_describeInferenceSchedulerCmd)
}

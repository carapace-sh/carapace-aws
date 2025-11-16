package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_startInferenceSchedulerCmd = &cobra.Command{
	Use:   "start-inference-scheduler",
	Short: "Starts an inference scheduler.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_startInferenceSchedulerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_startInferenceSchedulerCmd).Standalone()

		lookoutequipment_startInferenceSchedulerCmd.Flags().String("inference-scheduler-name", "", "The name of the inference scheduler to be started.")
		lookoutequipment_startInferenceSchedulerCmd.MarkFlagRequired("inference-scheduler-name")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_startInferenceSchedulerCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_stopInferenceSchedulerCmd = &cobra.Command{
	Use:   "stop-inference-scheduler",
	Short: "Stops an inference scheduler.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_stopInferenceSchedulerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_stopInferenceSchedulerCmd).Standalone()

		lookoutequipment_stopInferenceSchedulerCmd.Flags().String("inference-scheduler-name", "", "The name of the inference scheduler to be stopped.")
		lookoutequipment_stopInferenceSchedulerCmd.MarkFlagRequired("inference-scheduler-name")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_stopInferenceSchedulerCmd)
}

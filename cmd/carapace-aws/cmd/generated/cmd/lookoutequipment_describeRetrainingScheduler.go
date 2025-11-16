package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_describeRetrainingSchedulerCmd = &cobra.Command{
	Use:   "describe-retraining-scheduler",
	Short: "Provides a description of the retraining scheduler, including information such as the model name and retraining parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_describeRetrainingSchedulerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_describeRetrainingSchedulerCmd).Standalone()

		lookoutequipment_describeRetrainingSchedulerCmd.Flags().String("model-name", "", "The name of the model that the retraining scheduler is attached to.")
		lookoutequipment_describeRetrainingSchedulerCmd.MarkFlagRequired("model-name")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_describeRetrainingSchedulerCmd)
}

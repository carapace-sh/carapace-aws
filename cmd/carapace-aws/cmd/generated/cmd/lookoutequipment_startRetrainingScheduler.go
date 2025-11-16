package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_startRetrainingSchedulerCmd = &cobra.Command{
	Use:   "start-retraining-scheduler",
	Short: "Starts a retraining scheduler.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_startRetrainingSchedulerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_startRetrainingSchedulerCmd).Standalone()

		lookoutequipment_startRetrainingSchedulerCmd.Flags().String("model-name", "", "The name of the model whose retraining scheduler you want to start.")
		lookoutequipment_startRetrainingSchedulerCmd.MarkFlagRequired("model-name")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_startRetrainingSchedulerCmd)
}

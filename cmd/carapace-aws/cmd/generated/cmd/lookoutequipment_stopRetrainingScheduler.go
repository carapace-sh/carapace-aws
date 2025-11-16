package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_stopRetrainingSchedulerCmd = &cobra.Command{
	Use:   "stop-retraining-scheduler",
	Short: "Stops a retraining scheduler.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_stopRetrainingSchedulerCmd).Standalone()

	lookoutequipment_stopRetrainingSchedulerCmd.Flags().String("model-name", "", "The name of the model whose retraining scheduler you want to stop.")
	lookoutequipment_stopRetrainingSchedulerCmd.MarkFlagRequired("model-name")
	lookoutequipmentCmd.AddCommand(lookoutequipment_stopRetrainingSchedulerCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_deleteRetrainingSchedulerCmd = &cobra.Command{
	Use:   "delete-retraining-scheduler",
	Short: "Deletes a retraining scheduler from a model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_deleteRetrainingSchedulerCmd).Standalone()

	lookoutequipment_deleteRetrainingSchedulerCmd.Flags().String("model-name", "", "The name of the model whose retraining scheduler you want to delete.")
	lookoutequipment_deleteRetrainingSchedulerCmd.MarkFlagRequired("model-name")
	lookoutequipmentCmd.AddCommand(lookoutequipment_deleteRetrainingSchedulerCmd)
}

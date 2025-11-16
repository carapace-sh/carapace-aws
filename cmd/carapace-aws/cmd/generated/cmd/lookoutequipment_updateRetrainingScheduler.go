package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_updateRetrainingSchedulerCmd = &cobra.Command{
	Use:   "update-retraining-scheduler",
	Short: "Updates a retraining scheduler.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_updateRetrainingSchedulerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_updateRetrainingSchedulerCmd).Standalone()

		lookoutequipment_updateRetrainingSchedulerCmd.Flags().String("lookback-window", "", "The number of past days of data that will be used for retraining.")
		lookoutequipment_updateRetrainingSchedulerCmd.Flags().String("model-name", "", "The name of the model whose retraining scheduler you want to update.")
		lookoutequipment_updateRetrainingSchedulerCmd.Flags().String("promote-mode", "", "Indicates how the service will use new models.")
		lookoutequipment_updateRetrainingSchedulerCmd.Flags().String("retraining-frequency", "", "This parameter uses the [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#Durations) standard to set the frequency at which you want retraining to occur in terms of Years, Months, and/or Days (note: other parameters like Time are not currently supported).")
		lookoutequipment_updateRetrainingSchedulerCmd.Flags().String("retraining-start-date", "", "The start date for the retraining scheduler.")
		lookoutequipment_updateRetrainingSchedulerCmd.MarkFlagRequired("model-name")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_updateRetrainingSchedulerCmd)
}

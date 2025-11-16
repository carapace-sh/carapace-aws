package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_createRetrainingSchedulerCmd = &cobra.Command{
	Use:   "create-retraining-scheduler",
	Short: "Creates a retraining scheduler on the specified model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_createRetrainingSchedulerCmd).Standalone()

	lookoutequipment_createRetrainingSchedulerCmd.Flags().String("client-token", "", "A unique identifier for the request.")
	lookoutequipment_createRetrainingSchedulerCmd.Flags().String("lookback-window", "", "The number of past days of data that will be used for retraining.")
	lookoutequipment_createRetrainingSchedulerCmd.Flags().String("model-name", "", "The name of the model to add the retraining scheduler to.")
	lookoutequipment_createRetrainingSchedulerCmd.Flags().String("promote-mode", "", "Indicates how the service will use new models.")
	lookoutequipment_createRetrainingSchedulerCmd.Flags().String("retraining-frequency", "", "This parameter uses the [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#Durations) standard to set the frequency at which you want retraining to occur in terms of Years, Months, and/or Days (note: other parameters like Time are not currently supported).")
	lookoutequipment_createRetrainingSchedulerCmd.Flags().String("retraining-start-date", "", "The start date for the retraining scheduler.")
	lookoutequipment_createRetrainingSchedulerCmd.MarkFlagRequired("client-token")
	lookoutequipment_createRetrainingSchedulerCmd.MarkFlagRequired("lookback-window")
	lookoutequipment_createRetrainingSchedulerCmd.MarkFlagRequired("model-name")
	lookoutequipment_createRetrainingSchedulerCmd.MarkFlagRequired("retraining-frequency")
	lookoutequipmentCmd.AddCommand(lookoutequipment_createRetrainingSchedulerCmd)
}

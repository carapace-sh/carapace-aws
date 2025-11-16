package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_listInferenceEventsCmd = &cobra.Command{
	Use:   "list-inference-events",
	Short: "Lists all inference events that have been found for the specified inference scheduler.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_listInferenceEventsCmd).Standalone()

	lookoutequipment_listInferenceEventsCmd.Flags().String("inference-scheduler-name", "", "The name of the inference scheduler for the inference events listed.")
	lookoutequipment_listInferenceEventsCmd.Flags().String("interval-end-time", "", "Returns all the inference events with an end start time equal to or greater than less than the end time given.")
	lookoutequipment_listInferenceEventsCmd.Flags().String("interval-start-time", "", "Lookout for Equipment will return all the inference events with an end time equal to or greater than the start time given.")
	lookoutequipment_listInferenceEventsCmd.Flags().String("max-results", "", "Specifies the maximum number of inference events to list.")
	lookoutequipment_listInferenceEventsCmd.Flags().String("next-token", "", "An opaque pagination token indicating where to continue the listing of inference events.")
	lookoutequipment_listInferenceEventsCmd.MarkFlagRequired("inference-scheduler-name")
	lookoutequipment_listInferenceEventsCmd.MarkFlagRequired("interval-end-time")
	lookoutequipment_listInferenceEventsCmd.MarkFlagRequired("interval-start-time")
	lookoutequipmentCmd.AddCommand(lookoutequipment_listInferenceEventsCmd)
}

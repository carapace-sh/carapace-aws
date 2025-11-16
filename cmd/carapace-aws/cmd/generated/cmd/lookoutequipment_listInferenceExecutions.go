package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_listInferenceExecutionsCmd = &cobra.Command{
	Use:   "list-inference-executions",
	Short: "Lists all inference executions that have been performed by the specified inference scheduler.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_listInferenceExecutionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_listInferenceExecutionsCmd).Standalone()

		lookoutequipment_listInferenceExecutionsCmd.Flags().String("data-end-time-before", "", "The time reference in the inferenced dataset before which Amazon Lookout for Equipment stopped the inference execution.")
		lookoutequipment_listInferenceExecutionsCmd.Flags().String("data-start-time-after", "", "The time reference in the inferenced dataset after which Amazon Lookout for Equipment started the inference execution.")
		lookoutequipment_listInferenceExecutionsCmd.Flags().String("inference-scheduler-name", "", "The name of the inference scheduler for the inference execution listed.")
		lookoutequipment_listInferenceExecutionsCmd.Flags().String("max-results", "", "Specifies the maximum number of inference executions to list.")
		lookoutequipment_listInferenceExecutionsCmd.Flags().String("next-token", "", "An opaque pagination token indicating where to continue the listing of inference executions.")
		lookoutequipment_listInferenceExecutionsCmd.Flags().String("status", "", "The status of the inference execution.")
		lookoutequipment_listInferenceExecutionsCmd.MarkFlagRequired("inference-scheduler-name")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_listInferenceExecutionsCmd)
}

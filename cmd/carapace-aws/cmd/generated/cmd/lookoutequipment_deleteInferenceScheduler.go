package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_deleteInferenceSchedulerCmd = &cobra.Command{
	Use:   "delete-inference-scheduler",
	Short: "Deletes an inference scheduler that has been set up.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_deleteInferenceSchedulerCmd).Standalone()

	lookoutequipment_deleteInferenceSchedulerCmd.Flags().String("inference-scheduler-name", "", "The name of the inference scheduler to be deleted.")
	lookoutequipment_deleteInferenceSchedulerCmd.MarkFlagRequired("inference-scheduler-name")
	lookoutequipmentCmd.AddCommand(lookoutequipment_deleteInferenceSchedulerCmd)
}

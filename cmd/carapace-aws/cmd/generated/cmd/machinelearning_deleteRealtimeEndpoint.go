package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_deleteRealtimeEndpointCmd = &cobra.Command{
	Use:   "delete-realtime-endpoint",
	Short: "Deletes a real time endpoint of an `MLModel`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_deleteRealtimeEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(machinelearning_deleteRealtimeEndpointCmd).Standalone()

		machinelearning_deleteRealtimeEndpointCmd.Flags().String("mlmodel-id", "", "The ID assigned to the `MLModel` during creation.")
		machinelearning_deleteRealtimeEndpointCmd.MarkFlagRequired("mlmodel-id")
	})
	machinelearningCmd.AddCommand(machinelearning_deleteRealtimeEndpointCmd)
}

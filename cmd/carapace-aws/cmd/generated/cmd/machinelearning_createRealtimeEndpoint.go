package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_createRealtimeEndpointCmd = &cobra.Command{
	Use:   "create-realtime-endpoint",
	Short: "Creates a real-time endpoint for the `MLModel`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_createRealtimeEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(machinelearning_createRealtimeEndpointCmd).Standalone()

		machinelearning_createRealtimeEndpointCmd.Flags().String("mlmodel-id", "", "The ID assigned to the `MLModel` during creation.")
		machinelearning_createRealtimeEndpointCmd.MarkFlagRequired("mlmodel-id")
	})
	machinelearningCmd.AddCommand(machinelearning_createRealtimeEndpointCmd)
}

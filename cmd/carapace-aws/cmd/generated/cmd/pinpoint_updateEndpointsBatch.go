package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateEndpointsBatchCmd = &cobra.Command{
	Use:   "update-endpoints-batch",
	Short: "Creates a new batch of endpoints for an application or updates the settings and attributes of a batch of existing endpoints for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateEndpointsBatchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_updateEndpointsBatchCmd).Standalone()

		pinpoint_updateEndpointsBatchCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_updateEndpointsBatchCmd.Flags().String("endpoint-batch-request", "", "")
		pinpoint_updateEndpointsBatchCmd.MarkFlagRequired("application-id")
		pinpoint_updateEndpointsBatchCmd.MarkFlagRequired("endpoint-batch-request")
	})
	pinpointCmd.AddCommand(pinpoint_updateEndpointsBatchCmd)
}

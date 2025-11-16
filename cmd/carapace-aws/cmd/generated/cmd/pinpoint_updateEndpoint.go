package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateEndpointCmd = &cobra.Command{
	Use:   "update-endpoint",
	Short: "Creates a new endpoint for an application or updates the settings and attributes of an existing endpoint for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_updateEndpointCmd).Standalone()

		pinpoint_updateEndpointCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_updateEndpointCmd.Flags().String("endpoint-id", "", "The case insensitive unique identifier for the endpoint.")
		pinpoint_updateEndpointCmd.Flags().String("endpoint-request", "", "")
		pinpoint_updateEndpointCmd.MarkFlagRequired("application-id")
		pinpoint_updateEndpointCmd.MarkFlagRequired("endpoint-id")
		pinpoint_updateEndpointCmd.MarkFlagRequired("endpoint-request")
	})
	pinpointCmd.AddCommand(pinpoint_updateEndpointCmd)
}

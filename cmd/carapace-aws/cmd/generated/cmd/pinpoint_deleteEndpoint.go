package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteEndpointCmd = &cobra.Command{
	Use:   "delete-endpoint",
	Short: "Deletes an endpoint from an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_deleteEndpointCmd).Standalone()

		pinpoint_deleteEndpointCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_deleteEndpointCmd.Flags().String("endpoint-id", "", "The case insensitive unique identifier for the endpoint.")
		pinpoint_deleteEndpointCmd.MarkFlagRequired("application-id")
		pinpoint_deleteEndpointCmd.MarkFlagRequired("endpoint-id")
	})
	pinpointCmd.AddCommand(pinpoint_deleteEndpointCmd)
}

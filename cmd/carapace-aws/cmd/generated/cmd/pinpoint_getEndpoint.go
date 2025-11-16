package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getEndpointCmd = &cobra.Command{
	Use:   "get-endpoint",
	Short: "Retrieves information about the settings and attributes of a specific endpoint for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getEndpointCmd).Standalone()

	pinpoint_getEndpointCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_getEndpointCmd.Flags().String("endpoint-id", "", "The case insensitive unique identifier for the endpoint.")
	pinpoint_getEndpointCmd.MarkFlagRequired("application-id")
	pinpoint_getEndpointCmd.MarkFlagRequired("endpoint-id")
	pinpointCmd.AddCommand(pinpoint_getEndpointCmd)
}

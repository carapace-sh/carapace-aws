package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getUserEndpointsCmd = &cobra.Command{
	Use:   "get-user-endpoints",
	Short: "Retrieves information about all the endpoints that are associated with a specific user ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getUserEndpointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getUserEndpointsCmd).Standalone()

		pinpoint_getUserEndpointsCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_getUserEndpointsCmd.Flags().String("user-id", "", "The unique identifier for the user.")
		pinpoint_getUserEndpointsCmd.MarkFlagRequired("application-id")
		pinpoint_getUserEndpointsCmd.MarkFlagRequired("user-id")
	})
	pinpointCmd.AddCommand(pinpoint_getUserEndpointsCmd)
}

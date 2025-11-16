package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteUserEndpointsCmd = &cobra.Command{
	Use:   "delete-user-endpoints",
	Short: "Deletes all the endpoints that are associated with a specific user ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteUserEndpointsCmd).Standalone()

	pinpoint_deleteUserEndpointsCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_deleteUserEndpointsCmd.Flags().String("user-id", "", "The unique identifier for the user.")
	pinpoint_deleteUserEndpointsCmd.MarkFlagRequired("application-id")
	pinpoint_deleteUserEndpointsCmd.MarkFlagRequired("user-id")
	pinpointCmd.AddCommand(pinpoint_deleteUserEndpointsCmd)
}

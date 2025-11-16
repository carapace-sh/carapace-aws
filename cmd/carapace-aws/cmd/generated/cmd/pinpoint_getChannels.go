package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getChannelsCmd = &cobra.Command{
	Use:   "get-channels",
	Short: "Retrieves information about the history and status of each channel for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getChannelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getChannelsCmd).Standalone()

		pinpoint_getChannelsCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_getChannelsCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_getChannelsCmd)
}

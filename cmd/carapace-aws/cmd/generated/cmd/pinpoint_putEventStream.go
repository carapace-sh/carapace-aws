package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_putEventStreamCmd = &cobra.Command{
	Use:   "put-event-stream",
	Short: "Creates a new event stream for an application or updates the settings of an existing event stream for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_putEventStreamCmd).Standalone()

	pinpoint_putEventStreamCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_putEventStreamCmd.Flags().String("write-event-stream", "", "")
	pinpoint_putEventStreamCmd.MarkFlagRequired("application-id")
	pinpoint_putEventStreamCmd.MarkFlagRequired("write-event-stream")
	pinpointCmd.AddCommand(pinpoint_putEventStreamCmd)
}

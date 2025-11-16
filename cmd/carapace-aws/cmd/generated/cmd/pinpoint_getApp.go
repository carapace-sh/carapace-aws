package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getAppCmd = &cobra.Command{
	Use:   "get-app",
	Short: "Retrieves information about an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getAppCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getAppCmd).Standalone()

		pinpoint_getAppCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_getAppCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_getAppCmd)
}

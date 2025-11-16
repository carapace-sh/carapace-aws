package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteAppCmd = &cobra.Command{
	Use:   "delete-app",
	Short: "Deletes an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteAppCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_deleteAppCmd).Standalone()

		pinpoint_deleteAppCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_deleteAppCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_deleteAppCmd)
}

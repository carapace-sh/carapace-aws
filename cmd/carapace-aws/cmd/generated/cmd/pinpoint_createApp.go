package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_createAppCmd = &cobra.Command{
	Use:   "create-app",
	Short: "Creates an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_createAppCmd).Standalone()

	pinpoint_createAppCmd.Flags().String("create-application-request", "", "")
	pinpoint_createAppCmd.MarkFlagRequired("create-application-request")
	pinpointCmd.AddCommand(pinpoint_createAppCmd)
}

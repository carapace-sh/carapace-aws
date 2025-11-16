package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudcontrol_getResourceRequestStatusCmd = &cobra.Command{
	Use:   "get-resource-request-status",
	Short: "Returns the current status of a resource operation request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudcontrol_getResourceRequestStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudcontrol_getResourceRequestStatusCmd).Standalone()

		cloudcontrol_getResourceRequestStatusCmd.Flags().String("request-token", "", "A unique token used to track the progress of the resource operation request.")
		cloudcontrol_getResourceRequestStatusCmd.MarkFlagRequired("request-token")
	})
	cloudcontrolCmd.AddCommand(cloudcontrol_getResourceRequestStatusCmd)
}

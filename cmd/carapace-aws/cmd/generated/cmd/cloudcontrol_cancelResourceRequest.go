package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudcontrol_cancelResourceRequestCmd = &cobra.Command{
	Use:   "cancel-resource-request",
	Short: "Cancels the specified resource operation request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudcontrol_cancelResourceRequestCmd).Standalone()

	cloudcontrol_cancelResourceRequestCmd.Flags().String("request-token", "", "The `RequestToken` of the `ProgressEvent` object returned by the resource operation request.")
	cloudcontrol_cancelResourceRequestCmd.MarkFlagRequired("request-token")
	cloudcontrolCmd.AddCommand(cloudcontrol_cancelResourceRequestCmd)
}

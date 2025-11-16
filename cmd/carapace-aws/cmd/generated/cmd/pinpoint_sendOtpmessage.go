package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_sendOtpmessageCmd = &cobra.Command{
	Use:   "send-otpmessage",
	Short: "Send an OTP message",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_sendOtpmessageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_sendOtpmessageCmd).Standalone()

		pinpoint_sendOtpmessageCmd.Flags().String("application-id", "", "The unique ID of your Amazon Pinpoint application.")
		pinpoint_sendOtpmessageCmd.Flags().String("send-otpmessage-request-parameters", "", "")
		pinpoint_sendOtpmessageCmd.MarkFlagRequired("application-id")
		pinpoint_sendOtpmessageCmd.MarkFlagRequired("send-otpmessage-request-parameters")
	})
	pinpointCmd.AddCommand(pinpoint_sendOtpmessageCmd)
}

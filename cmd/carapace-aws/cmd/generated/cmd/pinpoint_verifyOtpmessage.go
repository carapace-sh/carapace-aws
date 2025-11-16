package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_verifyOtpmessageCmd = &cobra.Command{
	Use:   "verify-otpmessage",
	Short: "Verify an OTP",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_verifyOtpmessageCmd).Standalone()

	pinpoint_verifyOtpmessageCmd.Flags().String("application-id", "", "The unique ID of your Amazon Pinpoint application.")
	pinpoint_verifyOtpmessageCmd.Flags().String("verify-otpmessage-request-parameters", "", "")
	pinpoint_verifyOtpmessageCmd.MarkFlagRequired("application-id")
	pinpoint_verifyOtpmessageCmd.MarkFlagRequired("verify-otpmessage-request-parameters")
	pinpointCmd.AddCommand(pinpoint_verifyOtpmessageCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_carrierLookupCmd = &cobra.Command{
	Use:   "carrier-lookup",
	Short: "Returns information about a destination phone number, including whether the number type and whether it is valid, the carrier, and more.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_carrierLookupCmd).Standalone()

	pinpointSmsVoiceV2_carrierLookupCmd.Flags().String("phone-number", "", "The phone number that you want to retrieve information about.")
	pinpointSmsVoiceV2_carrierLookupCmd.MarkFlagRequired("phone-number")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_carrierLookupCmd)
}

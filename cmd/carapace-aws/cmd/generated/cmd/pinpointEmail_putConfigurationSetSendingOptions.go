package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_putConfigurationSetSendingOptionsCmd = &cobra.Command{
	Use:   "put-configuration-set-sending-options",
	Short: "Enable or disable email sending for messages that use a particular configuration set in a specific AWS Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_putConfigurationSetSendingOptionsCmd).Standalone()

	pinpointEmail_putConfigurationSetSendingOptionsCmd.Flags().String("configuration-set-name", "", "The name of the configuration set that you want to enable or disable email sending for.")
	pinpointEmail_putConfigurationSetSendingOptionsCmd.Flags().String("sending-enabled", "", "If `true`, email sending is enabled for the configuration set.")
	pinpointEmail_putConfigurationSetSendingOptionsCmd.MarkFlagRequired("configuration-set-name")
	pinpointEmailCmd.AddCommand(pinpointEmail_putConfigurationSetSendingOptionsCmd)
}

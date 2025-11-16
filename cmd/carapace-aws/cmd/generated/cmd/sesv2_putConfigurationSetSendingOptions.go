package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putConfigurationSetSendingOptionsCmd = &cobra.Command{
	Use:   "put-configuration-set-sending-options",
	Short: "Enable or disable email sending for messages that use a particular configuration set in a specific Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putConfigurationSetSendingOptionsCmd).Standalone()

	sesv2_putConfigurationSetSendingOptionsCmd.Flags().String("configuration-set-name", "", "The name of the configuration set to enable or disable email sending for.")
	sesv2_putConfigurationSetSendingOptionsCmd.Flags().String("sending-enabled", "", "If `true`, email sending is enabled for the configuration set.")
	sesv2_putConfigurationSetSendingOptionsCmd.MarkFlagRequired("configuration-set-name")
	sesv2Cmd.AddCommand(sesv2_putConfigurationSetSendingOptionsCmd)
}

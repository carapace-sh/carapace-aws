package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_updateConfigurationSetSendingEnabledCmd = &cobra.Command{
	Use:   "update-configuration-set-sending-enabled",
	Short: "Enables or disables email sending for messages sent using a specific configuration set in a given Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_updateConfigurationSetSendingEnabledCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_updateConfigurationSetSendingEnabledCmd).Standalone()

		ses_updateConfigurationSetSendingEnabledCmd.Flags().String("configuration-set-name", "", "The name of the configuration set to update.")
		ses_updateConfigurationSetSendingEnabledCmd.Flags().String("enabled", "", "Describes whether email sending is enabled or disabled for the configuration set.")
		ses_updateConfigurationSetSendingEnabledCmd.MarkFlagRequired("configuration-set-name")
		ses_updateConfigurationSetSendingEnabledCmd.MarkFlagRequired("enabled")
	})
	sesCmd.AddCommand(ses_updateConfigurationSetSendingEnabledCmd)
}

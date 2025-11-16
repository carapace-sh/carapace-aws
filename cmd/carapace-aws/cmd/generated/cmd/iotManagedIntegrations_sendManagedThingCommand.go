package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_sendManagedThingCommandCmd = &cobra.Command{
	Use:   "send-managed-thing-command",
	Short: "Send the command to the device represented by the managed thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_sendManagedThingCommandCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_sendManagedThingCommandCmd).Standalone()

		iotManagedIntegrations_sendManagedThingCommandCmd.Flags().String("account-association-id", "", "The identifier of the account association to use when sending a command to a managed thing.")
		iotManagedIntegrations_sendManagedThingCommandCmd.Flags().String("connector-association-id", "", "The ID tracking the current discovery process for one connector association.")
		iotManagedIntegrations_sendManagedThingCommandCmd.Flags().String("endpoints", "", "The device endpoint.")
		iotManagedIntegrations_sendManagedThingCommandCmd.Flags().String("managed-thing-id", "", "The id of the device.")
		iotManagedIntegrations_sendManagedThingCommandCmd.MarkFlagRequired("endpoints")
		iotManagedIntegrations_sendManagedThingCommandCmd.MarkFlagRequired("managed-thing-id")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_sendManagedThingCommandCmd)
}

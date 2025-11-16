package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_updateEventConfigurationByResourceTypesCmd = &cobra.Command{
	Use:   "update-event-configuration-by-resource-types",
	Short: "Update the event configuration based on resource types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_updateEventConfigurationByResourceTypesCmd).Standalone()

	iotwireless_updateEventConfigurationByResourceTypesCmd.Flags().String("connection-status", "", "Connection status resource type event configuration object for enabling and disabling wireless gateway topic.")
	iotwireless_updateEventConfigurationByResourceTypesCmd.Flags().String("device-registration-state", "", "Device registration state resource type event configuration object for enabling and disabling wireless gateway topic.")
	iotwireless_updateEventConfigurationByResourceTypesCmd.Flags().String("join", "", "Join resource type event configuration object for enabling and disabling wireless device topic.")
	iotwireless_updateEventConfigurationByResourceTypesCmd.Flags().String("message-delivery-status", "", "Message delivery status resource type event configuration object for enabling and disabling wireless device topic.")
	iotwireless_updateEventConfigurationByResourceTypesCmd.Flags().String("proximity", "", "Proximity resource type event configuration object for enabling and disabling wireless gateway topic.")
	iotwirelessCmd.AddCommand(iotwireless_updateEventConfigurationByResourceTypesCmd)
}

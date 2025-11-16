package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_updateResourceEventConfigurationCmd = &cobra.Command{
	Use:   "update-resource-event-configuration",
	Short: "Update the event configuration for a particular resource identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_updateResourceEventConfigurationCmd).Standalone()

	iotwireless_updateResourceEventConfigurationCmd.Flags().String("connection-status", "", "Event configuration for the connection status event.")
	iotwireless_updateResourceEventConfigurationCmd.Flags().String("device-registration-state", "", "Event configuration for the device registration state event.")
	iotwireless_updateResourceEventConfigurationCmd.Flags().String("identifier", "", "Resource identifier to opt in for event messaging.")
	iotwireless_updateResourceEventConfigurationCmd.Flags().String("identifier-type", "", "Identifier type of the particular resource identifier for event configuration.")
	iotwireless_updateResourceEventConfigurationCmd.Flags().String("join", "", "Event configuration for the join event.")
	iotwireless_updateResourceEventConfigurationCmd.Flags().String("message-delivery-status", "", "Event configuration for the message delivery status event.")
	iotwireless_updateResourceEventConfigurationCmd.Flags().String("partner-type", "", "Partner type of the resource if the identifier type is `PartnerAccountId`")
	iotwireless_updateResourceEventConfigurationCmd.Flags().String("proximity", "", "Event configuration for the proximity event.")
	iotwireless_updateResourceEventConfigurationCmd.MarkFlagRequired("identifier")
	iotwireless_updateResourceEventConfigurationCmd.MarkFlagRequired("identifier-type")
	iotwirelessCmd.AddCommand(iotwireless_updateResourceEventConfigurationCmd)
}

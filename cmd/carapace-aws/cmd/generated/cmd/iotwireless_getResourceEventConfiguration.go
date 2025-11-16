package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getResourceEventConfigurationCmd = &cobra.Command{
	Use:   "get-resource-event-configuration",
	Short: "Get the event configuration for a particular resource identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getResourceEventConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_getResourceEventConfigurationCmd).Standalone()

		iotwireless_getResourceEventConfigurationCmd.Flags().String("identifier", "", "Resource identifier to opt in for event messaging.")
		iotwireless_getResourceEventConfigurationCmd.Flags().String("identifier-type", "", "Identifier type of the particular resource identifier for event configuration.")
		iotwireless_getResourceEventConfigurationCmd.Flags().String("partner-type", "", "Partner type of the resource if the identifier type is `PartnerAccountId`.")
		iotwireless_getResourceEventConfigurationCmd.MarkFlagRequired("identifier")
		iotwireless_getResourceEventConfigurationCmd.MarkFlagRequired("identifier-type")
	})
	iotwirelessCmd.AddCommand(iotwireless_getResourceEventConfigurationCmd)
}

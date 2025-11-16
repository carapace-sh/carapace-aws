package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getEventConfigurationByResourceTypesCmd = &cobra.Command{
	Use:   "get-event-configuration-by-resource-types",
	Short: "Get the event configuration based on resource types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getEventConfigurationByResourceTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_getEventConfigurationByResourceTypesCmd).Standalone()

	})
	iotwirelessCmd.AddCommand(iotwireless_getEventConfigurationByResourceTypesCmd)
}

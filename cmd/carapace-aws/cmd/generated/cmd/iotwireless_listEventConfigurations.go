package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_listEventConfigurationsCmd = &cobra.Command{
	Use:   "list-event-configurations",
	Short: "List event configurations where at least one event topic has been enabled.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_listEventConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_listEventConfigurationsCmd).Standalone()

		iotwireless_listEventConfigurationsCmd.Flags().String("max-results", "", "")
		iotwireless_listEventConfigurationsCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
		iotwireless_listEventConfigurationsCmd.Flags().String("resource-type", "", "Resource type to filter event configurations.")
		iotwireless_listEventConfigurationsCmd.MarkFlagRequired("resource-type")
	})
	iotwirelessCmd.AddCommand(iotwireless_listEventConfigurationsCmd)
}

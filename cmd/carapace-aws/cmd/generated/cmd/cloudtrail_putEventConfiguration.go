package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_putEventConfigurationCmd = &cobra.Command{
	Use:   "put-event-configuration",
	Short: "Updates the event configuration settings for the specified event data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_putEventConfigurationCmd).Standalone()

	cloudtrail_putEventConfigurationCmd.Flags().String("context-key-selectors", "", "A list of context key selectors that will be included to provide enriched event data.")
	cloudtrail_putEventConfigurationCmd.Flags().String("event-data-store", "", "The Amazon Resource Name (ARN) or ID suffix of the ARN of the event data store for which you want to update event configuration settings.")
	cloudtrail_putEventConfigurationCmd.Flags().String("max-event-size", "", "The maximum allowed size for events to be stored in the specified event data store.")
	cloudtrail_putEventConfigurationCmd.MarkFlagRequired("context-key-selectors")
	cloudtrail_putEventConfigurationCmd.MarkFlagRequired("max-event-size")
	cloudtrailCmd.AddCommand(cloudtrail_putEventConfigurationCmd)
}

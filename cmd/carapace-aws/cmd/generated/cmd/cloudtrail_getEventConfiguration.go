package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_getEventConfigurationCmd = &cobra.Command{
	Use:   "get-event-configuration",
	Short: "Retrieves the current event configuration settings for the specified event data store, including details about maximum event size and context key selectors configured for the event data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_getEventConfigurationCmd).Standalone()

	cloudtrail_getEventConfigurationCmd.Flags().String("event-data-store", "", "The Amazon Resource Name (ARN) or ID suffix of the ARN of the event data store for which you want to retrieve event configuration settings.")
	cloudtrailCmd.AddCommand(cloudtrail_getEventConfigurationCmd)
}

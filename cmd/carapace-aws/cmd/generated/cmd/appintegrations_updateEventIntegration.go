package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_updateEventIntegrationCmd = &cobra.Command{
	Use:   "update-event-integration",
	Short: "Updates the description of an event integration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_updateEventIntegrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appintegrations_updateEventIntegrationCmd).Standalone()

		appintegrations_updateEventIntegrationCmd.Flags().String("description", "", "The description of the event integration.")
		appintegrations_updateEventIntegrationCmd.Flags().String("name", "", "The name of the event integration.")
		appintegrations_updateEventIntegrationCmd.MarkFlagRequired("name")
	})
	appintegrationsCmd.AddCommand(appintegrations_updateEventIntegrationCmd)
}

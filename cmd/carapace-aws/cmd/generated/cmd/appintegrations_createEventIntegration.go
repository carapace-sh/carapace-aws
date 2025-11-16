package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_createEventIntegrationCmd = &cobra.Command{
	Use:   "create-event-integration",
	Short: "Creates an EventIntegration, given a specified name, description, and a reference to an Amazon EventBridge bus in your account and a partner event source that pushes events to that bus.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_createEventIntegrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appintegrations_createEventIntegrationCmd).Standalone()

		appintegrations_createEventIntegrationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		appintegrations_createEventIntegrationCmd.Flags().String("description", "", "The description of the event integration.")
		appintegrations_createEventIntegrationCmd.Flags().String("event-bridge-bus", "", "The EventBridge bus.")
		appintegrations_createEventIntegrationCmd.Flags().String("event-filter", "", "The event filter.")
		appintegrations_createEventIntegrationCmd.Flags().String("name", "", "The name of the event integration.")
		appintegrations_createEventIntegrationCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		appintegrations_createEventIntegrationCmd.MarkFlagRequired("event-bridge-bus")
		appintegrations_createEventIntegrationCmd.MarkFlagRequired("event-filter")
		appintegrations_createEventIntegrationCmd.MarkFlagRequired("name")
	})
	appintegrationsCmd.AddCommand(appintegrations_createEventIntegrationCmd)
}

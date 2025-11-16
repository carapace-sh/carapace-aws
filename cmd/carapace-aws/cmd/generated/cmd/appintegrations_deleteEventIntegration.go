package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_deleteEventIntegrationCmd = &cobra.Command{
	Use:   "delete-event-integration",
	Short: "Deletes the specified existing event integration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_deleteEventIntegrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appintegrations_deleteEventIntegrationCmd).Standalone()

		appintegrations_deleteEventIntegrationCmd.Flags().String("name", "", "The name of the event integration.")
		appintegrations_deleteEventIntegrationCmd.MarkFlagRequired("name")
	})
	appintegrationsCmd.AddCommand(appintegrations_deleteEventIntegrationCmd)
}

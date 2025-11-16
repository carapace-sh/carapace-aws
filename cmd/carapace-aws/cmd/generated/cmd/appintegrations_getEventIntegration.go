package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_getEventIntegrationCmd = &cobra.Command{
	Use:   "get-event-integration",
	Short: "Returns information about the event integration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_getEventIntegrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appintegrations_getEventIntegrationCmd).Standalone()

		appintegrations_getEventIntegrationCmd.Flags().String("name", "", "The name of the event integration.")
		appintegrations_getEventIntegrationCmd.MarkFlagRequired("name")
	})
	appintegrationsCmd.AddCommand(appintegrations_getEventIntegrationCmd)
}

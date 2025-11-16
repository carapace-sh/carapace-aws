package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_listNotificationConfigurationsCmd = &cobra.Command{
	Use:   "list-notification-configurations",
	Short: "List all notification configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_listNotificationConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_listNotificationConfigurationsCmd).Standalone()

		iotManagedIntegrations_listNotificationConfigurationsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iotManagedIntegrations_listNotificationConfigurationsCmd.Flags().String("next-token", "", "A token that can be used to retrieve the next set of results.")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_listNotificationConfigurationsCmd)
}

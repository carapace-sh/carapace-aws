package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_listEventLogConfigurationsCmd = &cobra.Command{
	Use:   "list-event-log-configurations",
	Short: "List all event log configurations for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_listEventLogConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_listEventLogConfigurationsCmd).Standalone()

		iotManagedIntegrations_listEventLogConfigurationsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iotManagedIntegrations_listEventLogConfigurationsCmd.Flags().String("next-token", "", "A token that can be used to retrieve the next set of results.")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_listEventLogConfigurationsCmd)
}

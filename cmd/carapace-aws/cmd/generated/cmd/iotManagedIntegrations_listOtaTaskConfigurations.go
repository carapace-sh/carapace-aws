package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_listOtaTaskConfigurationsCmd = &cobra.Command{
	Use:   "list-ota-task-configurations",
	Short: "List all of the over-the-air (OTA) task configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_listOtaTaskConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_listOtaTaskConfigurationsCmd).Standalone()

		iotManagedIntegrations_listOtaTaskConfigurationsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iotManagedIntegrations_listOtaTaskConfigurationsCmd.Flags().String("next-token", "", "A token that can be used to retrieve the next set of results.")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_listOtaTaskConfigurationsCmd)
}

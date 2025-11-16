package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_listOtaTasksCmd = &cobra.Command{
	Use:   "list-ota-tasks",
	Short: "List all of the over-the-air (OTA) tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_listOtaTasksCmd).Standalone()

	iotManagedIntegrations_listOtaTasksCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iotManagedIntegrations_listOtaTasksCmd.Flags().String("next-token", "", "A token that can be used to retrieve the next set of results.")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_listOtaTasksCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_listOtaTaskExecutionsCmd = &cobra.Command{
	Use:   "list-ota-task-executions",
	Short: "List all of the over-the-air (OTA) task executions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_listOtaTaskExecutionsCmd).Standalone()

	iotManagedIntegrations_listOtaTaskExecutionsCmd.Flags().String("identifier", "", "The over-the-air (OTA) task id.")
	iotManagedIntegrations_listOtaTaskExecutionsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iotManagedIntegrations_listOtaTaskExecutionsCmd.Flags().String("next-token", "", "A token that can be used to retrieve the next set of results.")
	iotManagedIntegrations_listOtaTaskExecutionsCmd.MarkFlagRequired("identifier")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_listOtaTaskExecutionsCmd)
}

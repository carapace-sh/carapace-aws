package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_deleteDestinationCmd = &cobra.Command{
	Use:   "delete-destination",
	Short: "Deletes a notification destination specified by name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_deleteDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_deleteDestinationCmd).Standalone()

		iotManagedIntegrations_deleteDestinationCmd.Flags().String("name", "", "The id of the customer-managed destination.")
		iotManagedIntegrations_deleteDestinationCmd.MarkFlagRequired("name")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_deleteDestinationCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_listManagedThingsCmd = &cobra.Command{
	Use:   "list-managed-things",
	Short: "Listing all managed things with provision for filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_listManagedThingsCmd).Standalone()

	iotManagedIntegrations_listManagedThingsCmd.Flags().String("connector-destination-id-filter", "", "Filter managed things by the connector destination ID they are associated with.")
	iotManagedIntegrations_listManagedThingsCmd.Flags().String("connector-device-id-filter", "", "Filter managed things by the connector device ID they are associated with.")
	iotManagedIntegrations_listManagedThingsCmd.Flags().String("connector-policy-id-filter", "", "Filter on a connector policy id for a managed thing.")
	iotManagedIntegrations_listManagedThingsCmd.Flags().String("credential-locker-filter", "", "Filter on a credential locker for a managed thing.")
	iotManagedIntegrations_listManagedThingsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iotManagedIntegrations_listManagedThingsCmd.Flags().String("next-token", "", "A token that can be used to retrieve the next set of results.")
	iotManagedIntegrations_listManagedThingsCmd.Flags().String("owner-filter", "", "Filter on device owners when listing managed things.")
	iotManagedIntegrations_listManagedThingsCmd.Flags().String("parent-controller-identifier-filter", "", "Filter on a parent controller id for a managed thing.")
	iotManagedIntegrations_listManagedThingsCmd.Flags().String("provisioning-status-filter", "", "Filter on the status of the device.")
	iotManagedIntegrations_listManagedThingsCmd.Flags().String("role-filter", "", "Filter on the type of device used.")
	iotManagedIntegrations_listManagedThingsCmd.Flags().String("serial-number-filter", "", "Filter on the serial number of the device.")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_listManagedThingsCmd)
}

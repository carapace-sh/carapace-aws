package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_listManagedThingSchemasCmd = &cobra.Command{
	Use:   "list-managed-thing-schemas",
	Short: "List schemas associated with a managed thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_listManagedThingSchemasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_listManagedThingSchemasCmd).Standalone()

		iotManagedIntegrations_listManagedThingSchemasCmd.Flags().String("capability-id-filter", "", "Filter on a capability id.")
		iotManagedIntegrations_listManagedThingSchemasCmd.Flags().String("endpoint-id-filter", "", "Filter on an endpoint id.")
		iotManagedIntegrations_listManagedThingSchemasCmd.Flags().String("identifier", "", "The managed thing id.")
		iotManagedIntegrations_listManagedThingSchemasCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iotManagedIntegrations_listManagedThingSchemasCmd.Flags().String("next-token", "", "A token that can be used to retrieve the next set of results.")
		iotManagedIntegrations_listManagedThingSchemasCmd.MarkFlagRequired("identifier")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_listManagedThingSchemasCmd)
}

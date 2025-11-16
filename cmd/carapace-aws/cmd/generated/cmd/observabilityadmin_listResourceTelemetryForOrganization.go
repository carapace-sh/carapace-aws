package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_listResourceTelemetryForOrganizationCmd = &cobra.Command{
	Use:   "list-resource-telemetry-for-organization",
	Short: "Returns a list of telemetry configurations for Amazon Web Services resources supported by telemetry config in the organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_listResourceTelemetryForOrganizationCmd).Standalone()

	observabilityadmin_listResourceTelemetryForOrganizationCmd.Flags().String("account-identifiers", "", "A list of Amazon Web Services accounts used to filter the resources to those associated with the specified accounts.")
	observabilityadmin_listResourceTelemetryForOrganizationCmd.Flags().String("max-results", "", "A number field used to limit the number of results within the returned list.")
	observabilityadmin_listResourceTelemetryForOrganizationCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	observabilityadmin_listResourceTelemetryForOrganizationCmd.Flags().String("resource-identifier-prefix", "", "A string used to filter resources in the organization which have a `ResourceIdentifier` starting with the `ResourceIdentifierPrefix`.")
	observabilityadmin_listResourceTelemetryForOrganizationCmd.Flags().String("resource-tags", "", "A key-value pair to filter resources in the organization based on tags associated with the resource.")
	observabilityadmin_listResourceTelemetryForOrganizationCmd.Flags().String("resource-types", "", "A list of resource types used to filter resources in the organization.")
	observabilityadmin_listResourceTelemetryForOrganizationCmd.Flags().String("telemetry-configuration-state", "", "A key-value pair to filter resources in the organization based on the telemetry type and the state of the telemetry configuration.")
	observabilityadminCmd.AddCommand(observabilityadmin_listResourceTelemetryForOrganizationCmd)
}

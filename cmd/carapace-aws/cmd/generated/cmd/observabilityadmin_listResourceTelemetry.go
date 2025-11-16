package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_listResourceTelemetryCmd = &cobra.Command{
	Use:   "list-resource-telemetry",
	Short: "Returns a list of telemetry configurations for Amazon Web Services resources supported by telemetry config.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_listResourceTelemetryCmd).Standalone()

	observabilityadmin_listResourceTelemetryCmd.Flags().String("max-results", "", "A number field used to limit the number of results within the returned list.")
	observabilityadmin_listResourceTelemetryCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	observabilityadmin_listResourceTelemetryCmd.Flags().String("resource-identifier-prefix", "", "A string used to filter resources which have a `ResourceIdentifier` starting with the `ResourceIdentifierPrefix`.")
	observabilityadmin_listResourceTelemetryCmd.Flags().String("resource-tags", "", "A key-value pair to filter resources based on tags associated with the resource.")
	observabilityadmin_listResourceTelemetryCmd.Flags().String("resource-types", "", "A list of resource types used to filter resources supported by telemetry config.")
	observabilityadmin_listResourceTelemetryCmd.Flags().String("telemetry-configuration-state", "", "A key-value pair to filter resources based on the telemetry type and the state of the telemetry configuration.")
	observabilityadminCmd.AddCommand(observabilityadmin_listResourceTelemetryCmd)
}

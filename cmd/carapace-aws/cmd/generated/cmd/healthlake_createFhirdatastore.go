package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var healthlake_createFhirdatastoreCmd = &cobra.Command{
	Use:   "create-fhirdatastore",
	Short: "Create a FHIR-enabled data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(healthlake_createFhirdatastoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(healthlake_createFhirdatastoreCmd).Standalone()

		healthlake_createFhirdatastoreCmd.Flags().String("client-token", "", "An optional user-provided token to ensure API idempotency.")
		healthlake_createFhirdatastoreCmd.Flags().String("datastore-name", "", "The data store name (user-generated).")
		healthlake_createFhirdatastoreCmd.Flags().String("datastore-type-version", "", "The FHIR release version supported by the data store.")
		healthlake_createFhirdatastoreCmd.Flags().String("identity-provider-configuration", "", "The identity provider configuration to use for the data store.")
		healthlake_createFhirdatastoreCmd.Flags().String("preload-data-config", "", "An optional parameter to preload (import) open source Synthea FHIR data upon creation of the data store.")
		healthlake_createFhirdatastoreCmd.Flags().String("sse-configuration", "", "The server-side encryption key configuration for a customer-provided encryption key specified for creating a data store.")
		healthlake_createFhirdatastoreCmd.Flags().String("tags", "", "The resource tags applied to a data store when it is created.")
		healthlake_createFhirdatastoreCmd.MarkFlagRequired("datastore-type-version")
	})
	healthlakeCmd.AddCommand(healthlake_createFhirdatastoreCmd)
}

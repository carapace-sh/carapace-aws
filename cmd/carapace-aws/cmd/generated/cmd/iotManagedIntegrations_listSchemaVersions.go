package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_listSchemaVersionsCmd = &cobra.Command{
	Use:   "list-schema-versions",
	Short: "Lists schema versions with the provided information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_listSchemaVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_listSchemaVersionsCmd).Standalone()

		iotManagedIntegrations_listSchemaVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iotManagedIntegrations_listSchemaVersionsCmd.Flags().String("namespace", "", "Filter on the name of the schema version.")
		iotManagedIntegrations_listSchemaVersionsCmd.Flags().String("next-token", "", "A token that can be used to retrieve the next set of results.")
		iotManagedIntegrations_listSchemaVersionsCmd.Flags().String("schema-id", "", "Filter on the id of the schema version.")
		iotManagedIntegrations_listSchemaVersionsCmd.Flags().String("semantic-version", "", "The schema version.")
		iotManagedIntegrations_listSchemaVersionsCmd.Flags().String("type", "", "Filter on the type of schema version.")
		iotManagedIntegrations_listSchemaVersionsCmd.Flags().String("visibility", "", "The visibility of the schema version.")
		iotManagedIntegrations_listSchemaVersionsCmd.MarkFlagRequired("type")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_listSchemaVersionsCmd)
}

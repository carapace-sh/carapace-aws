package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getSchemaVersionCmd = &cobra.Command{
	Use:   "get-schema-version",
	Short: "Gets a schema version with the provided information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getSchemaVersionCmd).Standalone()

	iotManagedIntegrations_getSchemaVersionCmd.Flags().String("format", "", "The format of the schema version.")
	iotManagedIntegrations_getSchemaVersionCmd.Flags().String("schema-versioned-id", "", "Schema id with a version specified.")
	iotManagedIntegrations_getSchemaVersionCmd.Flags().String("type", "", "The type of schema version.")
	iotManagedIntegrations_getSchemaVersionCmd.MarkFlagRequired("schema-versioned-id")
	iotManagedIntegrations_getSchemaVersionCmd.MarkFlagRequired("type")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getSchemaVersionCmd)
}

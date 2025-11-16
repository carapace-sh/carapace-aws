package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_getDiscoveredSchemaCmd = &cobra.Command{
	Use:   "get-discovered-schema",
	Short: "Get the discovered schema that was generated based on sampled events.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_getDiscoveredSchemaCmd).Standalone()

	schemas_getDiscoveredSchemaCmd.Flags().String("events", "", "An array of strings where each string is a JSON event.")
	schemas_getDiscoveredSchemaCmd.Flags().String("type", "", "The type of event.")
	schemas_getDiscoveredSchemaCmd.MarkFlagRequired("events")
	schemas_getDiscoveredSchemaCmd.MarkFlagRequired("type")
	schemasCmd.AddCommand(schemas_getDiscoveredSchemaCmd)
}

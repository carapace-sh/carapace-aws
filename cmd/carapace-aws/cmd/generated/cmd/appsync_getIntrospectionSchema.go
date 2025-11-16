package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_getIntrospectionSchemaCmd = &cobra.Command{
	Use:   "get-introspection-schema",
	Short: "Retrieves the introspection schema for a GraphQL API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_getIntrospectionSchemaCmd).Standalone()

	appsync_getIntrospectionSchemaCmd.Flags().String("api-id", "", "The API ID.")
	appsync_getIntrospectionSchemaCmd.Flags().String("format", "", "The schema format: SDL or JSON.")
	appsync_getIntrospectionSchemaCmd.Flags().String("include-directives", "", "A flag that specifies whether the schema introspection should contain directives.")
	appsync_getIntrospectionSchemaCmd.MarkFlagRequired("api-id")
	appsync_getIntrospectionSchemaCmd.MarkFlagRequired("format")
	appsyncCmd.AddCommand(appsync_getIntrospectionSchemaCmd)
}

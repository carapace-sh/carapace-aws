package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_startSchemaCreationCmd = &cobra.Command{
	Use:   "start-schema-creation",
	Short: "Adds a new schema to your GraphQL API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_startSchemaCreationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_startSchemaCreationCmd).Standalone()

		appsync_startSchemaCreationCmd.Flags().String("api-id", "", "The API ID.")
		appsync_startSchemaCreationCmd.Flags().String("definition", "", "The schema definition, in GraphQL schema language format.")
		appsync_startSchemaCreationCmd.MarkFlagRequired("api-id")
		appsync_startSchemaCreationCmd.MarkFlagRequired("definition")
	})
	appsyncCmd.AddCommand(appsync_startSchemaCreationCmd)
}

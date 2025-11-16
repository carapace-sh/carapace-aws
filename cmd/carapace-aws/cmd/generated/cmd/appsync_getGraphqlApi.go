package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_getGraphqlApiCmd = &cobra.Command{
	Use:   "get-graphql-api",
	Short: "Retrieves a `GraphqlApi` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_getGraphqlApiCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_getGraphqlApiCmd).Standalone()

		appsync_getGraphqlApiCmd.Flags().String("api-id", "", "The API ID for the GraphQL API.")
		appsync_getGraphqlApiCmd.MarkFlagRequired("api-id")
	})
	appsyncCmd.AddCommand(appsync_getGraphqlApiCmd)
}

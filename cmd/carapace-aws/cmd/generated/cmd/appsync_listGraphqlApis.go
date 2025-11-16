package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_listGraphqlApisCmd = &cobra.Command{
	Use:   "list-graphql-apis",
	Short: "Lists your GraphQL APIs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_listGraphqlApisCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_listGraphqlApisCmd).Standalone()

		appsync_listGraphqlApisCmd.Flags().String("api-type", "", "The value that indicates whether the GraphQL API is a standard API (`GRAPHQL`) or merged API (`MERGED`).")
		appsync_listGraphqlApisCmd.Flags().String("max-results", "", "The maximum number of results that you want the request to return.")
		appsync_listGraphqlApisCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which you can use to return the next set of items in the list.")
		appsync_listGraphqlApisCmd.Flags().String("owner", "", "The account owner of the GraphQL API.")
	})
	appsyncCmd.AddCommand(appsync_listGraphqlApisCmd)
}

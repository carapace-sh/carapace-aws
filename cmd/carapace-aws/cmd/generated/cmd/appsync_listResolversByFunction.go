package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_listResolversByFunctionCmd = &cobra.Command{
	Use:   "list-resolvers-by-function",
	Short: "List the resolvers that are associated with a specific function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_listResolversByFunctionCmd).Standalone()

	appsync_listResolversByFunctionCmd.Flags().String("api-id", "", "The API ID.")
	appsync_listResolversByFunctionCmd.Flags().String("function-id", "", "The function ID.")
	appsync_listResolversByFunctionCmd.Flags().String("max-results", "", "The maximum number of results that you want the request to return.")
	appsync_listResolversByFunctionCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which you can use to return the next set of items in the list.")
	appsync_listResolversByFunctionCmd.MarkFlagRequired("api-id")
	appsync_listResolversByFunctionCmd.MarkFlagRequired("function-id")
	appsyncCmd.AddCommand(appsync_listResolversByFunctionCmd)
}

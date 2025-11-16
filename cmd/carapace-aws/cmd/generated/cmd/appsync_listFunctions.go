package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_listFunctionsCmd = &cobra.Command{
	Use:   "list-functions",
	Short: "List multiple functions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_listFunctionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_listFunctionsCmd).Standalone()

		appsync_listFunctionsCmd.Flags().String("api-id", "", "The GraphQL API ID.")
		appsync_listFunctionsCmd.Flags().String("max-results", "", "The maximum number of results that you want the request to return.")
		appsync_listFunctionsCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which you can use to return the next set of items in the list.")
		appsync_listFunctionsCmd.MarkFlagRequired("api-id")
	})
	appsyncCmd.AddCommand(appsync_listFunctionsCmd)
}

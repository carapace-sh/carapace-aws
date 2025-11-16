package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_listResolversCmd = &cobra.Command{
	Use:   "list-resolvers",
	Short: "Lists the resolvers for a given API and type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_listResolversCmd).Standalone()

	appsync_listResolversCmd.Flags().String("api-id", "", "The API ID.")
	appsync_listResolversCmd.Flags().String("max-results", "", "The maximum number of results that you want the request to return.")
	appsync_listResolversCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which you can use to return the next set of items in the list.")
	appsync_listResolversCmd.Flags().String("type-name", "", "The type name.")
	appsync_listResolversCmd.MarkFlagRequired("api-id")
	appsync_listResolversCmd.MarkFlagRequired("type-name")
	appsyncCmd.AddCommand(appsync_listResolversCmd)
}

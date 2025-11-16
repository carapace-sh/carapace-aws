package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_deleteGraphqlApiCmd = &cobra.Command{
	Use:   "delete-graphql-api",
	Short: "Deletes a `GraphqlApi` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_deleteGraphqlApiCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_deleteGraphqlApiCmd).Standalone()

		appsync_deleteGraphqlApiCmd.Flags().String("api-id", "", "The API ID.")
		appsync_deleteGraphqlApiCmd.MarkFlagRequired("api-id")
	})
	appsyncCmd.AddCommand(appsync_deleteGraphqlApiCmd)
}

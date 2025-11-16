package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_listApiKeysCmd = &cobra.Command{
	Use:   "list-api-keys",
	Short: "Lists the API keys for a given API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_listApiKeysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_listApiKeysCmd).Standalone()

		appsync_listApiKeysCmd.Flags().String("api-id", "", "The API ID.")
		appsync_listApiKeysCmd.Flags().String("max-results", "", "The maximum number of results that you want the request to return.")
		appsync_listApiKeysCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which you can use to return the next set of items in the list.")
		appsync_listApiKeysCmd.MarkFlagRequired("api-id")
	})
	appsyncCmd.AddCommand(appsync_listApiKeysCmd)
}

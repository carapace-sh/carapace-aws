package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_listApisCmd = &cobra.Command{
	Use:   "list-apis",
	Short: "Lists the APIs in your AppSync account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_listApisCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_listApisCmd).Standalone()

		appsync_listApisCmd.Flags().String("max-results", "", "The maximum number of results that you want the request to return.")
		appsync_listApisCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which you can use to return the next set of items in the list.")
	})
	appsyncCmd.AddCommand(appsync_listApisCmd)
}

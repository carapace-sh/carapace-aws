package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_listSourceApiAssociationsCmd = &cobra.Command{
	Use:   "list-source-api-associations",
	Short: "Lists the `SourceApiAssociationSummary` data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_listSourceApiAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_listSourceApiAssociationsCmd).Standalone()

		appsync_listSourceApiAssociationsCmd.Flags().String("api-id", "", "The API ID.")
		appsync_listSourceApiAssociationsCmd.Flags().String("max-results", "", "The maximum number of results that you want the request to return.")
		appsync_listSourceApiAssociationsCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which you can use to return the next set of items in the list.")
		appsync_listSourceApiAssociationsCmd.MarkFlagRequired("api-id")
	})
	appsyncCmd.AddCommand(appsync_listSourceApiAssociationsCmd)
}

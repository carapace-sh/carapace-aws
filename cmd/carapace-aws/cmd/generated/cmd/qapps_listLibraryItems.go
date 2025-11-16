package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_listLibraryItemsCmd = &cobra.Command{
	Use:   "list-library-items",
	Short: "Lists the library items for Amazon Q Apps that are published and available for users in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_listLibraryItemsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qapps_listLibraryItemsCmd).Standalone()

		qapps_listLibraryItemsCmd.Flags().String("category-id", "", "Optional category to filter the library items by.")
		qapps_listLibraryItemsCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
		qapps_listLibraryItemsCmd.Flags().String("limit", "", "The maximum number of library items to return in the response.")
		qapps_listLibraryItemsCmd.Flags().String("next-token", "", "The token to request the next page of results.")
		qapps_listLibraryItemsCmd.MarkFlagRequired("instance-id")
	})
	qappsCmd.AddCommand(qapps_listLibraryItemsCmd)
}

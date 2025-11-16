package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_associateLibraryItemReviewCmd = &cobra.Command{
	Use:   "associate-library-item-review",
	Short: "Associates a rating or review for a library item with the user submitting the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_associateLibraryItemReviewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qapps_associateLibraryItemReviewCmd).Standalone()

		qapps_associateLibraryItemReviewCmd.Flags().String("instance-id", "", "The unique identifier for the Amazon Q Business application environment instance.")
		qapps_associateLibraryItemReviewCmd.Flags().String("library-item-id", "", "The unique identifier of the library item to associate the review with.")
		qapps_associateLibraryItemReviewCmd.MarkFlagRequired("instance-id")
		qapps_associateLibraryItemReviewCmd.MarkFlagRequired("library-item-id")
	})
	qappsCmd.AddCommand(qapps_associateLibraryItemReviewCmd)
}

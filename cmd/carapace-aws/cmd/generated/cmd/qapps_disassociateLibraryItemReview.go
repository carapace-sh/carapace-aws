package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_disassociateLibraryItemReviewCmd = &cobra.Command{
	Use:   "disassociate-library-item-review",
	Short: "Removes a rating or review previously submitted by the user for a library item.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_disassociateLibraryItemReviewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qapps_disassociateLibraryItemReviewCmd).Standalone()

		qapps_disassociateLibraryItemReviewCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
		qapps_disassociateLibraryItemReviewCmd.Flags().String("library-item-id", "", "The unique identifier of the library item to remove the review from.")
		qapps_disassociateLibraryItemReviewCmd.MarkFlagRequired("instance-id")
		qapps_disassociateLibraryItemReviewCmd.MarkFlagRequired("library-item-id")
	})
	qappsCmd.AddCommand(qapps_disassociateLibraryItemReviewCmd)
}

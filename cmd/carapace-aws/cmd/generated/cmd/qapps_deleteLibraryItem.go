package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_deleteLibraryItemCmd = &cobra.Command{
	Use:   "delete-library-item",
	Short: "Deletes a library item for an Amazon Q App, removing it from the library so it can no longer be discovered or used by other users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_deleteLibraryItemCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qapps_deleteLibraryItemCmd).Standalone()

		qapps_deleteLibraryItemCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
		qapps_deleteLibraryItemCmd.Flags().String("library-item-id", "", "The unique identifier of the library item to delete.")
		qapps_deleteLibraryItemCmd.MarkFlagRequired("instance-id")
		qapps_deleteLibraryItemCmd.MarkFlagRequired("library-item-id")
	})
	qappsCmd.AddCommand(qapps_deleteLibraryItemCmd)
}

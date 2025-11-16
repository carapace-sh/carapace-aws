package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_updateLibraryItemCmd = &cobra.Command{
	Use:   "update-library-item",
	Short: "Updates the library item for an Amazon Q App.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_updateLibraryItemCmd).Standalone()

	qapps_updateLibraryItemCmd.Flags().String("categories", "", "The new categories to associate with the library item.")
	qapps_updateLibraryItemCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
	qapps_updateLibraryItemCmd.Flags().String("library-item-id", "", "The unique identifier of the library item to update.")
	qapps_updateLibraryItemCmd.Flags().String("status", "", "The new status to set for the library item, such as \"Published\" or \"Hidden\".")
	qapps_updateLibraryItemCmd.MarkFlagRequired("instance-id")
	qapps_updateLibraryItemCmd.MarkFlagRequired("library-item-id")
	qappsCmd.AddCommand(qapps_updateLibraryItemCmd)
}

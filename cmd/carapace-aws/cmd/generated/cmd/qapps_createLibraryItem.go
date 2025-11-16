package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_createLibraryItemCmd = &cobra.Command{
	Use:   "create-library-item",
	Short: "Creates a new library item for an Amazon Q App, allowing it to be discovered and used by other allowed users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_createLibraryItemCmd).Standalone()

	qapps_createLibraryItemCmd.Flags().String("app-id", "", "The unique identifier of the Amazon Q App to publish to the library.")
	qapps_createLibraryItemCmd.Flags().String("app-version", "", "The version of the Amazon Q App to publish to the library.")
	qapps_createLibraryItemCmd.Flags().String("categories", "", "The categories to associate with the library item for easier discovery.")
	qapps_createLibraryItemCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
	qapps_createLibraryItemCmd.MarkFlagRequired("app-id")
	qapps_createLibraryItemCmd.MarkFlagRequired("app-version")
	qapps_createLibraryItemCmd.MarkFlagRequired("categories")
	qapps_createLibraryItemCmd.MarkFlagRequired("instance-id")
	qappsCmd.AddCommand(qapps_createLibraryItemCmd)
}

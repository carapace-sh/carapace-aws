package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_getLibraryItemCmd = &cobra.Command{
	Use:   "get-library-item",
	Short: "Retrieves details about a library item for an Amazon Q App, including its metadata, categories, ratings, and usage statistics.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_getLibraryItemCmd).Standalone()

	qapps_getLibraryItemCmd.Flags().String("app-id", "", "The unique identifier of the Amazon Q App associated with the library item.")
	qapps_getLibraryItemCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
	qapps_getLibraryItemCmd.Flags().String("library-item-id", "", "The unique identifier of the library item to retrieve.")
	qapps_getLibraryItemCmd.MarkFlagRequired("instance-id")
	qapps_getLibraryItemCmd.MarkFlagRequired("library-item-id")
	qappsCmd.AddCommand(qapps_getLibraryItemCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_updateLibraryItemMetadataCmd = &cobra.Command{
	Use:   "update-library-item-metadata",
	Short: "Updates the verification status of a library item for an Amazon Q App.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_updateLibraryItemMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qapps_updateLibraryItemMetadataCmd).Standalone()

		qapps_updateLibraryItemMetadataCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
		qapps_updateLibraryItemMetadataCmd.Flags().Bool("is-verified", false, "The verification status of the library item")
		qapps_updateLibraryItemMetadataCmd.Flags().String("library-item-id", "", "The unique identifier of the updated library item.")
		qapps_updateLibraryItemMetadataCmd.Flags().Bool("no-is-verified", false, "The verification status of the library item")
		qapps_updateLibraryItemMetadataCmd.MarkFlagRequired("instance-id")
		qapps_updateLibraryItemMetadataCmd.MarkFlagRequired("library-item-id")
		qapps_updateLibraryItemMetadataCmd.Flag("no-is-verified").Hidden = true
	})
	qappsCmd.AddCommand(qapps_updateLibraryItemMetadataCmd)
}

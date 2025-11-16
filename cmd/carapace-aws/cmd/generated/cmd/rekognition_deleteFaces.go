package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_deleteFacesCmd = &cobra.Command{
	Use:   "delete-faces",
	Short: "Deletes faces from a collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_deleteFacesCmd).Standalone()

	rekognition_deleteFacesCmd.Flags().String("collection-id", "", "Collection from which to remove the specific faces.")
	rekognition_deleteFacesCmd.Flags().String("face-ids", "", "An array of face IDs to delete.")
	rekognition_deleteFacesCmd.MarkFlagRequired("collection-id")
	rekognition_deleteFacesCmd.MarkFlagRequired("face-ids")
	rekognitionCmd.AddCommand(rekognition_deleteFacesCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_deleteCollectionCmd = &cobra.Command{
	Use:   "delete-collection",
	Short: "Deletes the specified collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_deleteCollectionCmd).Standalone()

	rekognition_deleteCollectionCmd.Flags().String("collection-id", "", "ID of the collection to delete.")
	rekognition_deleteCollectionCmd.MarkFlagRequired("collection-id")
	rekognitionCmd.AddCommand(rekognition_deleteCollectionCmd)
}

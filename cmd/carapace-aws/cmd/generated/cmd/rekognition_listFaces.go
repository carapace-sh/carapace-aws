package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_listFacesCmd = &cobra.Command{
	Use:   "list-faces",
	Short: "Returns metadata for faces in the specified collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_listFacesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_listFacesCmd).Standalone()

		rekognition_listFacesCmd.Flags().String("collection-id", "", "ID of the collection from which to list the faces.")
		rekognition_listFacesCmd.Flags().String("face-ids", "", "An array of face IDs to filter results with when listing faces in a collection.")
		rekognition_listFacesCmd.Flags().String("max-results", "", "Maximum number of faces to return.")
		rekognition_listFacesCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there is more data to retrieve), Amazon Rekognition returns a pagination token in the response.")
		rekognition_listFacesCmd.Flags().String("user-id", "", "An array of user IDs to filter results with when listing faces in a collection.")
		rekognition_listFacesCmd.MarkFlagRequired("collection-id")
	})
	rekognitionCmd.AddCommand(rekognition_listFacesCmd)
}

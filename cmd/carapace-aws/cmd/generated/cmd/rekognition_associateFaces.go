package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_associateFacesCmd = &cobra.Command{
	Use:   "associate-faces",
	Short: "Associates one or more faces with an existing UserID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_associateFacesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_associateFacesCmd).Standalone()

		rekognition_associateFacesCmd.Flags().String("client-request-token", "", "Idempotent token used to identify the request to `AssociateFaces`.")
		rekognition_associateFacesCmd.Flags().String("collection-id", "", "The ID of an existing collection containing the UserID.")
		rekognition_associateFacesCmd.Flags().String("face-ids", "", "An array of FaceIDs to associate with the UserID.")
		rekognition_associateFacesCmd.Flags().String("user-id", "", "The ID for the existing UserID.")
		rekognition_associateFacesCmd.Flags().String("user-match-threshold", "", "An optional value specifying the minimum confidence in the UserID match to return.")
		rekognition_associateFacesCmd.MarkFlagRequired("collection-id")
		rekognition_associateFacesCmd.MarkFlagRequired("face-ids")
		rekognition_associateFacesCmd.MarkFlagRequired("user-id")
	})
	rekognitionCmd.AddCommand(rekognition_associateFacesCmd)
}

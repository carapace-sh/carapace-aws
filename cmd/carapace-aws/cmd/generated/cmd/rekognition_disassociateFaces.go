package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_disassociateFacesCmd = &cobra.Command{
	Use:   "disassociate-faces",
	Short: "Removes the association between a `Face` supplied in an array of `FaceIds` and the User.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_disassociateFacesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_disassociateFacesCmd).Standalone()

		rekognition_disassociateFacesCmd.Flags().String("client-request-token", "", "Idempotent token used to identify the request to `DisassociateFaces`.")
		rekognition_disassociateFacesCmd.Flags().String("collection-id", "", "The ID of an existing collection containing the UserID.")
		rekognition_disassociateFacesCmd.Flags().String("face-ids", "", "An array of face IDs to disassociate from the UserID.")
		rekognition_disassociateFacesCmd.Flags().String("user-id", "", "ID for the existing UserID.")
		rekognition_disassociateFacesCmd.MarkFlagRequired("collection-id")
		rekognition_disassociateFacesCmd.MarkFlagRequired("face-ids")
		rekognition_disassociateFacesCmd.MarkFlagRequired("user-id")
	})
	rekognitionCmd.AddCommand(rekognition_disassociateFacesCmd)
}

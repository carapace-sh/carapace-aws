package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_importDocumentCmd = &cobra.Command{
	Use:   "import-document",
	Short: "Uploads a file that can then be used either as a default in a `FileUploadCard` from Q App definition or as a file that is used inside a single Q App run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_importDocumentCmd).Standalone()

	qapps_importDocumentCmd.Flags().String("app-id", "", "The unique identifier of the Q App the file is associated with.")
	qapps_importDocumentCmd.Flags().String("card-id", "", "The unique identifier of the card the file is associated with.")
	qapps_importDocumentCmd.Flags().String("file-contents-base64", "", "The base64-encoded contents of the file to upload.")
	qapps_importDocumentCmd.Flags().String("file-name", "", "The name of the file being uploaded.")
	qapps_importDocumentCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
	qapps_importDocumentCmd.Flags().String("scope", "", "Whether the file is associated with a Q App definition or a specific Q App session.")
	qapps_importDocumentCmd.Flags().String("session-id", "", "The unique identifier of the Q App session the file is associated with, if applicable.")
	qapps_importDocumentCmd.MarkFlagRequired("app-id")
	qapps_importDocumentCmd.MarkFlagRequired("card-id")
	qapps_importDocumentCmd.MarkFlagRequired("file-contents-base64")
	qapps_importDocumentCmd.MarkFlagRequired("file-name")
	qapps_importDocumentCmd.MarkFlagRequired("instance-id")
	qapps_importDocumentCmd.MarkFlagRequired("scope")
	qappsCmd.AddCommand(qapps_importDocumentCmd)
}

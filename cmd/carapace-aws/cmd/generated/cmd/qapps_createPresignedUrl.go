package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_createPresignedUrlCmd = &cobra.Command{
	Use:   "create-presigned-url",
	Short: "Creates a presigned URL for an S3 POST operation to upload a file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_createPresignedUrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qapps_createPresignedUrlCmd).Standalone()

		qapps_createPresignedUrlCmd.Flags().String("app-id", "", "The unique identifier of the Q App the file is associated with.")
		qapps_createPresignedUrlCmd.Flags().String("card-id", "", "The unique identifier of the card the file is associated with.")
		qapps_createPresignedUrlCmd.Flags().String("file-contents-sha256", "", "The Base64-encoded SHA-256 digest of the contents of the file to be uploaded.")
		qapps_createPresignedUrlCmd.Flags().String("file-name", "", "The name of the file to be uploaded.")
		qapps_createPresignedUrlCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
		qapps_createPresignedUrlCmd.Flags().String("scope", "", "Whether the file is associated with a Q App definition or a specific Q App session.")
		qapps_createPresignedUrlCmd.Flags().String("session-id", "", "The unique identifier of the Q App session the file is associated with, if applicable.")
		qapps_createPresignedUrlCmd.MarkFlagRequired("app-id")
		qapps_createPresignedUrlCmd.MarkFlagRequired("card-id")
		qapps_createPresignedUrlCmd.MarkFlagRequired("file-contents-sha256")
		qapps_createPresignedUrlCmd.MarkFlagRequired("file-name")
		qapps_createPresignedUrlCmd.MarkFlagRequired("instance-id")
		qapps_createPresignedUrlCmd.MarkFlagRequired("scope")
	})
	qappsCmd.AddCommand(qapps_createPresignedUrlCmd)
}

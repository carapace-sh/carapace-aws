package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_startContentUploadCmd = &cobra.Command{
	Use:   "start-content-upload",
	Short: "Get a URL to upload content to a knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_startContentUploadCmd).Standalone()

	qconnect_startContentUploadCmd.Flags().String("content-type", "", "The type of content to upload.")
	qconnect_startContentUploadCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	qconnect_startContentUploadCmd.Flags().String("presigned-url-time-to-live", "", "The expected expiration time of the generated presigned URL, specified in minutes.")
	qconnect_startContentUploadCmd.MarkFlagRequired("content-type")
	qconnect_startContentUploadCmd.MarkFlagRequired("knowledge-base-id")
	qconnectCmd.AddCommand(qconnect_startContentUploadCmd)
}

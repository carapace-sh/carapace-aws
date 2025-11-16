package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_startContentUploadCmd = &cobra.Command{
	Use:   "start-content-upload",
	Short: "Get a URL to upload content to a knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_startContentUploadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wisdom_startContentUploadCmd).Standalone()

		wisdom_startContentUploadCmd.Flags().String("content-type", "", "The type of content to upload.")
		wisdom_startContentUploadCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		wisdom_startContentUploadCmd.Flags().String("presigned-url-time-to-live", "", "The expected expiration time of the generated presigned URL, specified in minutes.")
		wisdom_startContentUploadCmd.MarkFlagRequired("content-type")
		wisdom_startContentUploadCmd.MarkFlagRequired("knowledge-base-id")
	})
	wisdomCmd.AddCommand(wisdom_startContentUploadCmd)
}

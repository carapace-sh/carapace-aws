package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_getMediaCmd = &cobra.Command{
	Use:   "get-media",
	Short: "Returns the image bytes corresponding to a media object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_getMediaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_getMediaCmd).Standalone()

		qbusiness_getMediaCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business which contains the media object.")
		qbusiness_getMediaCmd.Flags().String("conversation-id", "", "The identifier of the Amazon Q Business conversation.")
		qbusiness_getMediaCmd.Flags().String("media-id", "", "The identifier of the media object.")
		qbusiness_getMediaCmd.Flags().String("message-id", "", "The identifier of the Amazon Q Business message.")
		qbusiness_getMediaCmd.MarkFlagRequired("application-id")
		qbusiness_getMediaCmd.MarkFlagRequired("conversation-id")
		qbusiness_getMediaCmd.MarkFlagRequired("media-id")
		qbusiness_getMediaCmd.MarkFlagRequired("message-id")
	})
	qbusinessCmd.AddCommand(qbusiness_getMediaCmd)
}

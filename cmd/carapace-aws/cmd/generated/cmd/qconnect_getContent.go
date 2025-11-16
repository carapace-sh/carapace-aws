package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_getContentCmd = &cobra.Command{
	Use:   "get-content",
	Short: "Retrieves content, including a pre-signed URL to download the content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_getContentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_getContentCmd).Standalone()

		qconnect_getContentCmd.Flags().String("content-id", "", "The identifier of the content.")
		qconnect_getContentCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		qconnect_getContentCmd.MarkFlagRequired("content-id")
		qconnect_getContentCmd.MarkFlagRequired("knowledge-base-id")
	})
	qconnectCmd.AddCommand(qconnect_getContentCmd)
}

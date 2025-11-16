package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_createContentCmd = &cobra.Command{
	Use:   "create-content",
	Short: "Creates Amazon Q in Connect content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_createContentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_createContentCmd).Standalone()

		qconnect_createContentCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		qconnect_createContentCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		qconnect_createContentCmd.Flags().String("metadata", "", "A key/value map to store attributes without affecting tagging or recommendations.")
		qconnect_createContentCmd.Flags().String("name", "", "The name of the content.")
		qconnect_createContentCmd.Flags().String("override-link-out-uri", "", "The URI you want to use for the article.")
		qconnect_createContentCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		qconnect_createContentCmd.Flags().String("title", "", "The title of the content.")
		qconnect_createContentCmd.Flags().String("upload-id", "", "A pointer to the uploaded asset.")
		qconnect_createContentCmd.MarkFlagRequired("knowledge-base-id")
		qconnect_createContentCmd.MarkFlagRequired("name")
		qconnect_createContentCmd.MarkFlagRequired("upload-id")
	})
	qconnectCmd.AddCommand(qconnect_createContentCmd)
}

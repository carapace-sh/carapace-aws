package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_createContentCmd = &cobra.Command{
	Use:   "create-content",
	Short: "Creates Wisdom content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_createContentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wisdom_createContentCmd).Standalone()

		wisdom_createContentCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		wisdom_createContentCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		wisdom_createContentCmd.Flags().String("metadata", "", "A key/value map to store attributes without affecting tagging or recommendations.")
		wisdom_createContentCmd.Flags().String("name", "", "The name of the content.")
		wisdom_createContentCmd.Flags().String("override-link-out-uri", "", "The URI you want to use for the article.")
		wisdom_createContentCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		wisdom_createContentCmd.Flags().String("title", "", "The title of the content.")
		wisdom_createContentCmd.Flags().String("upload-id", "", "A pointer to the uploaded asset.")
		wisdom_createContentCmd.MarkFlagRequired("knowledge-base-id")
		wisdom_createContentCmd.MarkFlagRequired("name")
		wisdom_createContentCmd.MarkFlagRequired("upload-id")
	})
	wisdomCmd.AddCommand(wisdom_createContentCmd)
}

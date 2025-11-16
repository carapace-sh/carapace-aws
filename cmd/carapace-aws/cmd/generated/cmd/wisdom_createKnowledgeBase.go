package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_createKnowledgeBaseCmd = &cobra.Command{
	Use:   "create-knowledge-base",
	Short: "Creates a knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_createKnowledgeBaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wisdom_createKnowledgeBaseCmd).Standalone()

		wisdom_createKnowledgeBaseCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		wisdom_createKnowledgeBaseCmd.Flags().String("description", "", "The description.")
		wisdom_createKnowledgeBaseCmd.Flags().String("knowledge-base-type", "", "The type of knowledge base.")
		wisdom_createKnowledgeBaseCmd.Flags().String("name", "", "The name of the knowledge base.")
		wisdom_createKnowledgeBaseCmd.Flags().String("rendering-configuration", "", "Information about how to render the content.")
		wisdom_createKnowledgeBaseCmd.Flags().String("server-side-encryption-configuration", "", "The configuration information for the customer managed key used for encryption.")
		wisdom_createKnowledgeBaseCmd.Flags().String("source-configuration", "", "The source of the knowledge base content.")
		wisdom_createKnowledgeBaseCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		wisdom_createKnowledgeBaseCmd.MarkFlagRequired("knowledge-base-type")
		wisdom_createKnowledgeBaseCmd.MarkFlagRequired("name")
	})
	wisdomCmd.AddCommand(wisdom_createKnowledgeBaseCmd)
}

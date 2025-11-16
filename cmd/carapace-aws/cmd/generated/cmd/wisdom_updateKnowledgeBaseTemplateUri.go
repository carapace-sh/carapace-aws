package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_updateKnowledgeBaseTemplateUriCmd = &cobra.Command{
	Use:   "update-knowledge-base-template-uri",
	Short: "Updates the template URI of a knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_updateKnowledgeBaseTemplateUriCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wisdom_updateKnowledgeBaseTemplateUriCmd).Standalone()

		wisdom_updateKnowledgeBaseTemplateUriCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		wisdom_updateKnowledgeBaseTemplateUriCmd.Flags().String("template-uri", "", "The template URI to update.")
		wisdom_updateKnowledgeBaseTemplateUriCmd.MarkFlagRequired("knowledge-base-id")
		wisdom_updateKnowledgeBaseTemplateUriCmd.MarkFlagRequired("template-uri")
	})
	wisdomCmd.AddCommand(wisdom_updateKnowledgeBaseTemplateUriCmd)
}

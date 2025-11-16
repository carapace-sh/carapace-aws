package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_removeKnowledgeBaseTemplateUriCmd = &cobra.Command{
	Use:   "remove-knowledge-base-template-uri",
	Short: "Removes a URI template from a knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_removeKnowledgeBaseTemplateUriCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wisdom_removeKnowledgeBaseTemplateUriCmd).Standalone()

		wisdom_removeKnowledgeBaseTemplateUriCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		wisdom_removeKnowledgeBaseTemplateUriCmd.MarkFlagRequired("knowledge-base-id")
	})
	wisdomCmd.AddCommand(wisdom_removeKnowledgeBaseTemplateUriCmd)
}

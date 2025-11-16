package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_removeKnowledgeBaseTemplateUriCmd = &cobra.Command{
	Use:   "remove-knowledge-base-template-uri",
	Short: "Removes a URI template from a knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_removeKnowledgeBaseTemplateUriCmd).Standalone()

	qconnect_removeKnowledgeBaseTemplateUriCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	qconnect_removeKnowledgeBaseTemplateUriCmd.MarkFlagRequired("knowledge-base-id")
	qconnectCmd.AddCommand(qconnect_removeKnowledgeBaseTemplateUriCmd)
}

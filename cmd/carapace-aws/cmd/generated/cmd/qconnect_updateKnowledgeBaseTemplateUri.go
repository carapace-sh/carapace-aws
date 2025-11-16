package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_updateKnowledgeBaseTemplateUriCmd = &cobra.Command{
	Use:   "update-knowledge-base-template-uri",
	Short: "Updates the template URI of a knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_updateKnowledgeBaseTemplateUriCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_updateKnowledgeBaseTemplateUriCmd).Standalone()

		qconnect_updateKnowledgeBaseTemplateUriCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		qconnect_updateKnowledgeBaseTemplateUriCmd.Flags().String("template-uri", "", "The template URI to update.")
		qconnect_updateKnowledgeBaseTemplateUriCmd.MarkFlagRequired("knowledge-base-id")
		qconnect_updateKnowledgeBaseTemplateUriCmd.MarkFlagRequired("template-uri")
	})
	qconnectCmd.AddCommand(qconnect_updateKnowledgeBaseTemplateUriCmd)
}

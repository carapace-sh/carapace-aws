package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_deleteKnowledgeBaseCmd = &cobra.Command{
	Use:   "delete-knowledge-base",
	Short: "Deletes the knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_deleteKnowledgeBaseCmd).Standalone()

	qconnect_deleteKnowledgeBaseCmd.Flags().String("knowledge-base-id", "", "The knowledge base to delete content from.")
	qconnect_deleteKnowledgeBaseCmd.MarkFlagRequired("knowledge-base-id")
	qconnectCmd.AddCommand(qconnect_deleteKnowledgeBaseCmd)
}

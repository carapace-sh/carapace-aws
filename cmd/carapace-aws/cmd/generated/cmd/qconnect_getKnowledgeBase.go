package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_getKnowledgeBaseCmd = &cobra.Command{
	Use:   "get-knowledge-base",
	Short: "Retrieves information about the knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_getKnowledgeBaseCmd).Standalone()

	qconnect_getKnowledgeBaseCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	qconnect_getKnowledgeBaseCmd.MarkFlagRequired("knowledge-base-id")
	qconnectCmd.AddCommand(qconnect_getKnowledgeBaseCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_deleteKnowledgeBaseCmd = &cobra.Command{
	Use:   "delete-knowledge-base",
	Short: "Deletes the knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_deleteKnowledgeBaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wisdom_deleteKnowledgeBaseCmd).Standalone()

		wisdom_deleteKnowledgeBaseCmd.Flags().String("knowledge-base-id", "", "The knowledge base to delete content from.")
		wisdom_deleteKnowledgeBaseCmd.MarkFlagRequired("knowledge-base-id")
	})
	wisdomCmd.AddCommand(wisdom_deleteKnowledgeBaseCmd)
}

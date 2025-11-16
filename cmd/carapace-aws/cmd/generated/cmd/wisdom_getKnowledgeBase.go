package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_getKnowledgeBaseCmd = &cobra.Command{
	Use:   "get-knowledge-base",
	Short: "Retrieves information about the knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_getKnowledgeBaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wisdom_getKnowledgeBaseCmd).Standalone()

		wisdom_getKnowledgeBaseCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		wisdom_getKnowledgeBaseCmd.MarkFlagRequired("knowledge-base-id")
	})
	wisdomCmd.AddCommand(wisdom_getKnowledgeBaseCmd)
}

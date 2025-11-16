package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_updateKnowledgeBaseCmd = &cobra.Command{
	Use:   "update-knowledge-base",
	Short: "Updates the configuration of a knowledge base with the fields that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_updateKnowledgeBaseCmd).Standalone()

	bedrockAgent_updateKnowledgeBaseCmd.Flags().String("description", "", "Specifies a new description for the knowledge base.")
	bedrockAgent_updateKnowledgeBaseCmd.Flags().String("knowledge-base-configuration", "", "Specifies the configuration for the embeddings model used for the knowledge base.")
	bedrockAgent_updateKnowledgeBaseCmd.Flags().String("knowledge-base-id", "", "The unique identifier of the knowledge base to update.")
	bedrockAgent_updateKnowledgeBaseCmd.Flags().String("name", "", "Specifies a new name for the knowledge base.")
	bedrockAgent_updateKnowledgeBaseCmd.Flags().String("role-arn", "", "Specifies a different Amazon Resource Name (ARN) of the IAM role with permissions to invoke API operations on the knowledge base.")
	bedrockAgent_updateKnowledgeBaseCmd.Flags().String("storage-configuration", "", "Specifies the configuration for the vector store used for the knowledge base.")
	bedrockAgent_updateKnowledgeBaseCmd.MarkFlagRequired("knowledge-base-configuration")
	bedrockAgent_updateKnowledgeBaseCmd.MarkFlagRequired("knowledge-base-id")
	bedrockAgent_updateKnowledgeBaseCmd.MarkFlagRequired("name")
	bedrockAgent_updateKnowledgeBaseCmd.MarkFlagRequired("role-arn")
	bedrockAgentCmd.AddCommand(bedrockAgent_updateKnowledgeBaseCmd)
}

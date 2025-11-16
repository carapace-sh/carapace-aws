package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_createKnowledgeBaseCmd = &cobra.Command{
	Use:   "create-knowledge-base",
	Short: "Creates a knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_createKnowledgeBaseCmd).Standalone()

	bedrockAgent_createKnowledgeBaseCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
	bedrockAgent_createKnowledgeBaseCmd.Flags().String("description", "", "A description of the knowledge base.")
	bedrockAgent_createKnowledgeBaseCmd.Flags().String("knowledge-base-configuration", "", "Contains details about the embeddings model used for the knowledge base.")
	bedrockAgent_createKnowledgeBaseCmd.Flags().String("name", "", "A name for the knowledge base.")
	bedrockAgent_createKnowledgeBaseCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role with permissions to invoke API operations on the knowledge base.")
	bedrockAgent_createKnowledgeBaseCmd.Flags().String("storage-configuration", "", "Contains details about the configuration of the vector database used for the knowledge base.")
	bedrockAgent_createKnowledgeBaseCmd.Flags().String("tags", "", "Specify the key-value pairs for the tags that you want to attach to your knowledge base in this object.")
	bedrockAgent_createKnowledgeBaseCmd.MarkFlagRequired("knowledge-base-configuration")
	bedrockAgent_createKnowledgeBaseCmd.MarkFlagRequired("name")
	bedrockAgent_createKnowledgeBaseCmd.MarkFlagRequired("role-arn")
	bedrockAgentCmd.AddCommand(bedrockAgent_createKnowledgeBaseCmd)
}

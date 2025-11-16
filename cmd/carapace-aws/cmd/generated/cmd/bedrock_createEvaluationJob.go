package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_createEvaluationJobCmd = &cobra.Command{
	Use:   "create-evaluation-job",
	Short: "Creates an evaluation job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_createEvaluationJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_createEvaluationJobCmd).Standalone()

		bedrock_createEvaluationJobCmd.Flags().String("application-type", "", "Specifies whether the evaluation job is for evaluating a model or evaluating a knowledge base (retrieval and response generation).")
		bedrock_createEvaluationJobCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
		bedrock_createEvaluationJobCmd.Flags().String("customer-encryption-key-id", "", "Specify your customer managed encryption key Amazon Resource Name (ARN) that will be used to encrypt your evaluation job.")
		bedrock_createEvaluationJobCmd.Flags().String("evaluation-config", "", "Contains the configuration details of either an automated or human-based evaluation job.")
		bedrock_createEvaluationJobCmd.Flags().String("inference-config", "", "Contains the configuration details of the inference model for the evaluation job.")
		bedrock_createEvaluationJobCmd.Flags().String("job-description", "", "A description of the evaluation job.")
		bedrock_createEvaluationJobCmd.Flags().String("job-name", "", "A name for the evaluation job.")
		bedrock_createEvaluationJobCmd.Flags().String("job-tags", "", "Tags to attach to the model evaluation job.")
		bedrock_createEvaluationJobCmd.Flags().String("output-data-config", "", "Contains the configuration details of the Amazon S3 bucket for storing the results of the evaluation job.")
		bedrock_createEvaluationJobCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM service role that Amazon Bedrock can assume to perform tasks on your behalf.")
		bedrock_createEvaluationJobCmd.MarkFlagRequired("evaluation-config")
		bedrock_createEvaluationJobCmd.MarkFlagRequired("inference-config")
		bedrock_createEvaluationJobCmd.MarkFlagRequired("job-name")
		bedrock_createEvaluationJobCmd.MarkFlagRequired("output-data-config")
		bedrock_createEvaluationJobCmd.MarkFlagRequired("role-arn")
	})
	bedrockCmd.AddCommand(bedrock_createEvaluationJobCmd)
}

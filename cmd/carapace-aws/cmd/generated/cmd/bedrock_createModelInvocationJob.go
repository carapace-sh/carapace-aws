package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_createModelInvocationJobCmd = &cobra.Command{
	Use:   "create-model-invocation-job",
	Short: "Creates a batch inference job to invoke a model on multiple prompts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_createModelInvocationJobCmd).Standalone()

	bedrock_createModelInvocationJobCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
	bedrock_createModelInvocationJobCmd.Flags().String("input-data-config", "", "Details about the location of the input to the batch inference job.")
	bedrock_createModelInvocationJobCmd.Flags().String("job-name", "", "A name to give the batch inference job.")
	bedrock_createModelInvocationJobCmd.Flags().String("model-id", "", "The unique identifier of the foundation model to use for the batch inference job.")
	bedrock_createModelInvocationJobCmd.Flags().String("output-data-config", "", "Details about the location of the output of the batch inference job.")
	bedrock_createModelInvocationJobCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the service role with permissions to carry out and manage batch inference.")
	bedrock_createModelInvocationJobCmd.Flags().String("tags", "", "Any tags to associate with the batch inference job.")
	bedrock_createModelInvocationJobCmd.Flags().String("timeout-duration-in-hours", "", "The number of hours after which to force the batch inference job to time out.")
	bedrock_createModelInvocationJobCmd.Flags().String("vpc-config", "", "The configuration of the Virtual Private Cloud (VPC) for the data in the batch inference job.")
	bedrock_createModelInvocationJobCmd.MarkFlagRequired("input-data-config")
	bedrock_createModelInvocationJobCmd.MarkFlagRequired("job-name")
	bedrock_createModelInvocationJobCmd.MarkFlagRequired("model-id")
	bedrock_createModelInvocationJobCmd.MarkFlagRequired("output-data-config")
	bedrock_createModelInvocationJobCmd.MarkFlagRequired("role-arn")
	bedrockCmd.AddCommand(bedrock_createModelInvocationJobCmd)
}

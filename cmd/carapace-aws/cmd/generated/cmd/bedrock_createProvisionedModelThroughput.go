package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_createProvisionedModelThroughputCmd = &cobra.Command{
	Use:   "create-provisioned-model-throughput",
	Short: "Creates dedicated throughput for a base or custom model with the model units and for the duration that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_createProvisionedModelThroughputCmd).Standalone()

	bedrock_createProvisionedModelThroughputCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
	bedrock_createProvisionedModelThroughputCmd.Flags().String("commitment-duration", "", "The commitment duration requested for the Provisioned Throughput.")
	bedrock_createProvisionedModelThroughputCmd.Flags().String("model-id", "", "The Amazon Resource Name (ARN) or name of the model to associate with this Provisioned Throughput.")
	bedrock_createProvisionedModelThroughputCmd.Flags().String("model-units", "", "Number of model units to allocate.")
	bedrock_createProvisionedModelThroughputCmd.Flags().String("provisioned-model-name", "", "The name for this Provisioned Throughput.")
	bedrock_createProvisionedModelThroughputCmd.Flags().String("tags", "", "Tags to associate with this Provisioned Throughput.")
	bedrock_createProvisionedModelThroughputCmd.MarkFlagRequired("model-id")
	bedrock_createProvisionedModelThroughputCmd.MarkFlagRequired("model-units")
	bedrock_createProvisionedModelThroughputCmd.MarkFlagRequired("provisioned-model-name")
	bedrockCmd.AddCommand(bedrock_createProvisionedModelThroughputCmd)
}

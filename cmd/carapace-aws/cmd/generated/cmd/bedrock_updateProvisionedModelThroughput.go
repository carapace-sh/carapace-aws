package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_updateProvisionedModelThroughputCmd = &cobra.Command{
	Use:   "update-provisioned-model-throughput",
	Short: "Updates the name or associated model for a Provisioned Throughput.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_updateProvisionedModelThroughputCmd).Standalone()

	bedrock_updateProvisionedModelThroughputCmd.Flags().String("desired-model-id", "", "The Amazon Resource Name (ARN) of the new model to associate with this Provisioned Throughput.")
	bedrock_updateProvisionedModelThroughputCmd.Flags().String("desired-provisioned-model-name", "", "The new name for this Provisioned Throughput.")
	bedrock_updateProvisionedModelThroughputCmd.Flags().String("provisioned-model-id", "", "The Amazon Resource Name (ARN) or name of the Provisioned Throughput to update.")
	bedrock_updateProvisionedModelThroughputCmd.MarkFlagRequired("provisioned-model-id")
	bedrockCmd.AddCommand(bedrock_updateProvisionedModelThroughputCmd)
}
